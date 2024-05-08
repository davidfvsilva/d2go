package nip

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

const (
	RuleResultFullMatch RuleResult = 1
	RuleResultPartial   RuleResult = 2
	RuleResultNoMatch   RuleResult = 3
)

var (
	fixedPropsRegexp = regexp.MustCompile(`(\[(type|quality|class|name|flag|color)]\s*(<=|<|>|>=|!=|==)\s*([a-zA-Z0-9]+))`)
	statsRegexp      = regexp.MustCompile(`\[(.*?)]`)
	maxQtyRegexp     = regexp.MustCompile(`(\[maxquantity]\s*(<=|<|>|>=|!=|==)\s*([0-9]+))`)
)

type Rule struct {
	RawLine       string // Original line, don't use it for evaluation
	Filename      string
	LineNumber    int
	Enabled       bool
	maxQuantity   int
	stage1        *vm.Program
	stage2        *vm.Program
	requiredStats []string
}

type RuleResult int
type Rules []Rule

func (r Rules) EvaluateAll(it data.Item) (Rule, RuleResult) {
	bestMatch := RuleResultNoMatch
	bestMatchingRule := Rule{}
	for _, rule := range r {
		if rule.Enabled {
			result, err := rule.Evaluate(it)
			if err != nil {
				continue
			}
			if result == RuleResultFullMatch {
				return rule, result
			}
			if result == RuleResultPartial {
				bestMatch = result
				bestMatchingRule = rule
			}
		}
	}

	return bestMatchingRule, bestMatch
}

var fixedPropsList = map[string]int{"type": 0, "quality": 0, "class": 0, "name": 0, "flag": 0, "color": 0}

func NewRule(rawRule string, filename string, lineNumber int) (Rule, error) {
	rule := sanitizeLine(rawRule)

	// Try to get the maxquantity value and purge it from the rule, we can not evaluate it
	maxQuantity := 0
	for _, prop := range maxQtyRegexp.FindAllStringSubmatch(rule, -1) {
		mxQty, err := strconv.Atoi(prop[3])
		if err != nil {
			return Rule{}, fmt.Errorf("error parsing maxquantity value %s: %w", prop[3], err)
		}
		maxQuantity = mxQty
		rule = strings.ReplaceAll(rule, prop[0], "")
	}

	// Sanitize again, just in case we messed up the rule while parsing maxquantity
	rule = sanitizeLine(rule)
	if rule == "" {
		return Rule{}, ErrEmptyRule
	}

	r := Rule{
		RawLine:     rawRule,
		Filename:    filename,
		LineNumber:  lineNumber,
		Enabled:     true,
		maxQuantity: maxQuantity,
	}

	for i, stg := range strings.Split(rule, "#") {
		line := strings.TrimSpace(stg)
		if line == "" {
			break
		}

		if i == 0 {
			line, err := replaceStringPropertiesInStage1(line)
			if err != nil {
				return Rule{}, err
			}

			line = strings.ReplaceAll(line, "[", "")
			line = strings.ReplaceAll(line, "]", "")
			program, err := expr.Compile(line, expr.Env(fixedPropsList))
			if err != nil {
				return Rule{}, fmt.Errorf("error compiling rule: %w", err)
			}
			r.stage1 = program
		}

		if i == 1 {
			r.requiredStats = getRequiredStatsForRule(line)

			statsMap := make(map[string]int)
			for _, prop := range r.requiredStats {
				statsMap[prop] = 0
			}

			line = strings.ReplaceAll(line, "[", "")
			line = strings.ReplaceAll(line, "]", "")
			program, err := expr.Compile(line, expr.Env(statsMap))
			if err != nil {
				return Rule{}, fmt.Errorf("error compiling rule: %w", err)
			}
			r.stage2 = program

			// We only care about first and second part, third one is maxquantity and was already parsed
			break
		}
	}

	return r, nil
}

func (r Rule) Evaluate(it data.Item) (RuleResult, error) {
	stage1Props := make(map[string]int)
	for prop := range fixedPropsList {
		switch prop {
		case "type":
			stage1Props["type"] = it.Type().ID
		case "quality":
			stage1Props["quality"] = int(it.Quality)
		case "class":
			stage1Props["class"] = int(it.Desc().Tier())
		case "name":
			stage1Props["name"] = it.ID
		case "flag":
			stage1Props["flag"] = map[bool]int{true: 1, false: 0}[it.Ethereal]
		case "color":
			// TODO: Not supported yet
		}
	}

	// Let's evaluate first stage
	stage1Result, err := expr.Run(r.stage1, stage1Props)
	if err != nil {
		return RuleResultNoMatch, fmt.Errorf("error evaluating rule: %w", err)
	}

	// If stage1 does not match, we can stop here, nothing else to match
	if !stage1Result.(bool) {
		return RuleResultNoMatch, nil
	}

	stage2Result := true
	if r.stage2 != nil {
		stage2Props := make(map[string]int)
		for _, statName := range r.requiredStats {
			statData, found := statAliases[statName]
			if !found {
				return RuleResultNoMatch, fmt.Errorf("property %s is not valid or not supported", statName)
			}

			// Exception for enhanceddefense, is not accurate
			if strings.EqualFold(statName, "enhanceddefense") {
				enhancedDefense := r.calculateEnhancedDefense(it)
				stage2Props[statName] = enhancedDefense
				continue
			}

			layer := 0
			if len(statData) > 1 {
				layer = statData[1]
			}
			itemStat, _ := it.FindStat(stat.ID(statData[0]), layer)
			stage2Props[statName] = itemStat.Value
		}

		res, err := expr.Run(r.stage2, stage2Props)
		if err != nil {
			return RuleResultNoMatch, fmt.Errorf("error evaluating rule: %w", err)
		}
		stage2Result = res.(bool)
	}

	// 100% rule match, we can return here
	if stage1Result.(bool) && stage2Result {
		return RuleResultFullMatch, nil
	}

	// If Stage 1 matches and the item is NOT identified, let's return a partial match, we can not be 100% sure if all the stats are matching
	if stage1Result.(bool) && !it.Identified {
		return RuleResultPartial, nil
	}
	return RuleResultNoMatch, nil
}

func replaceStringPropertiesInStage1(stage1 string) (string, error) {
	baseProperties := fixedPropsRegexp.FindAllStringSubmatch(stage1, -1)
	for _, prop := range baseProperties {
		replaceWith := ""
		switch prop[2] {
		case "type":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", item.ItemTypes[typeAliases[prop[4]]].ID))
		case "quality":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", qualityAliases[prop[4]]))
		case "class":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", classAliases[prop[4]]))
		case "name":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", item.GetIDByName(prop[4])))
		case "flag":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", 1))
		case "color":
			// TODO: Not supported yet
			return "", fmt.Errorf("property %s is not supported yet", prop[2])
		}

		if replaceWith != "" {
			stage1 = strings.ReplaceAll(stage1, prop[0], replaceWith)
		}
	}

	return stage1, nil
}

func getRequiredStatsForRule(line string) []string {
	statsList := make([]string, 0)
	for _, statName := range statsRegexp.FindAllStringSubmatch(line, -1) {
		statsList = append(statsList, statName[1])
	}
	return statsList
}

// MaxQuantity returns the maximum quantity of items that character can have, 0 means no limit
func (r Rule) MaxQuantity() int {
	return r.maxQuantity
}

func (r Rule) calculateEnhancedDefense(i data.Item) int {
	defenseStat, found := i.FindStat(stat.Defense, 0)

	// This is special case for weapons and other stuff without base Defense and with EnhancedDefense, we can return it directly
	// since EnhancedDefense it's a stat.
	if !found {
		if edStat, edStatFound := i.FindStat(stat.EnhancedDefense, 0); edStatFound {
			return edStat.Value
		}
	}

	// Make it only work for white base items
	if i.Quality != item.QualitySuperior {
		return 0
	}

	if !i.Type().IsType(item.TypeArmor) {
		return 0
	}

	itemTypeMaxDefense := i.Desc().MaxDefense

	if defenseStat.Value <= itemTypeMaxDefense {
		return 0
	}

	return int(float64(defenseStat.Value-itemTypeMaxDefense) / float64(itemTypeMaxDefense) * 100)
}
