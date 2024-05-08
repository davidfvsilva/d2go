package nip

import "github.com/hectorgimenez/d2go/pkg/data/item"

var typeAliases = map[string]string{
	"shield":          item.TypeShield,
	"armor":           item.TypeArmor,
	"gold":            item.TypeGold,
	"bowquiver":       item.TypeBowQuiver,
	"crossbowquiver":  item.TypeCrossbowQuiver,
	"playerbodypart":  item.TypePlayerBodyPart,
	"herb":            item.TypeHerb,
	"potion":          item.TypePotion,
	"ring":            item.TypeRing,
	"elixir":          item.TypeElixir,
	"amulet":          item.TypeAmulet,
	"charm":           item.TypeCharm,
	"boots":           item.TypeBoots,
	"gloves":          item.TypeGloves,
	"book":            item.TypeBook,
	"belt":            item.TypeBelt,
	"gem":             item.TypeGem,
	"torch":           item.TypeTorch,
	"scroll":          item.TypeScroll,
	"scepter":         item.TypeScepter,
	"wand":            item.TypeWand,
	"staff":           item.TypeStaff,
	"bow":             item.TypeBow,
	"axe":             item.TypeAxe,
	"club":            item.TypeClub,
	"sword":           item.TypeSword,
	"hammer":          item.TypeHammer,
	"knife":           item.TypeKnife,
	"spear":           item.TypeSpear,
	"polearm":         item.TypePolearm,
	"crossbow":        item.TypeCrossbow,
	"mace":            item.TypeMace,
	"helm":            item.TypeHelm,
	"missilepotion":   item.TypeMissilePotion,
	"quest":           item.TypeQuest,
	"bodypart":        item.TypeBodyPart,
	"key":             item.TypeKey,
	"throwingknife":   item.TypeThrowingKnife,
	"throwingaxe":     item.TypeThrowingAxe,
	"javelin":         item.TypeJavelin,
	"weapon":          item.TypeWeapon,
	"meleeweapon":     item.TypeMeleeWeapon,
	"missileweapon":   item.TypeMissileWeapon,
	"thrownweapon":    item.TypeThrownWeapon,
	"comboweapon":     item.TypeComboWeapon,
	"anyarmor":        item.TypeAnyArmor,
	"anyshield":       item.TypeAnyShield,
	"miscellaneous":   item.TypeMiscellaneous,
	"socketfiller":    item.TypeSocketFiller,
	"secondhand":      item.TypeSecondHand,
	"stavesandrods":   item.TypeStavesAndRods,
	"missile":         item.TypeMissile,
	"blunt":           item.TypeBlunt,
	"jewel":           item.TypeJewel,
	"classspecific":   item.TypeClassSpecific,
	"amazonitem":      item.TypeAmazonItem,
	"barbarianitem":   item.TypeBarbarianItem,
	"necromanceritem": item.TypeNecromancerItem,
	"paladinitem":     item.TypePaladinItem,
	"sorceressitem":   item.TypeSorceressItem,
	"assassinitem":    item.TypeAssassinItem,
	"druiditem":       item.TypeDruidItem,
	"handtohand":      item.TypeHandtoHand,
	"orb":             item.TypeOrb,
	"voodooheads":     item.TypeVoodooHeads,
	"auricshields":    item.TypeAuricShields,
	"primalhelm":      item.TypePrimalHelm,
	"pelt":            item.TypePelt,
	"cloak":           item.TypeCloak,
	"rune":            item.TypeRune,
	"circlet":         item.TypeCirclet,
	"healingpotion":   item.TypeHealingPotion,
	"manapotion":      item.TypeManaPotion,
	"rejuvpotion":     item.TypeRejuvPotion,
	"staminapotion":   item.TypeStaminaPotion,
	"antidotepotion":  item.TypeAntidotePotion,
	"thawingpotion":   item.TypeThawingPotion,
	"smallcharm":      item.TypeSmallCharm,
	"mediumcharm":     item.TypeMediumCharm,
	"largecharm":      item.TypeLargeCharm,
	"amazonbow":       item.TypeAmazonBow,
	"amazonspear":     item.TypeAmazonSpear,
	"amazonjavelin":   item.TypeAmazonJavelin,
	"assassinclaw":    item.TypeHandtoHand2,
	"magicbowquiv":    item.TypeMagicBowQuiv,
	"magicxbowquiv":   item.TypeMagicXbowQuiv,
	"chippedgem":      item.TypeChippedGem,
	"flawedgem":       item.TypeFlawedGem,
	"standardgem":     item.TypeStandardGem,
	"flawlessgem":     item.TypeFlawlessGem,
	"perfectgem":      item.TypePerfectGem,
	"amethyst":        item.TypeAmethyst,
	"diamond":         item.TypeDiamond,
	"emerald":         item.TypeEmerald,
	"ruby":            item.TypeRuby,
	"sapphire":        item.TypeSapphire,
	"topaz":           item.TypeTopaz,
	"skull":           item.TypeSkull,
}

var qualityAliases = map[string]int{
	"lowquality": 1,
	"normal":     2,
	"superior":   3,
	"magic":      4,
	"set":        5,
	"rare":       6,
	"unique":     7,
	"crafted":    8,
}

var classAliases = map[string]int{
	"normal":      0,
	"exceptional": 1,
	"elite":       2,
}

var statAliases = map[string][]int{
	"strength":         {0},
	"energy":           {1},
	"dexterity":        {2},
	"vitality":         {3},
	"statpts":          {4},
	"newskills":        {5},
	"hitpoints":        {6},
	"maxhp":            {7},
	"mana":             {8},
	"maxmana":          {9},
	"stamina":          {10},
	"maxstamina":       {11},
	"level":            {12},
	"experience":       {13},
	"gold":             {14},
	"goldbank":         {15},
	"itemarmorpercent": {16, 0}, "enhanceddefense": {16, 0},
	"itemmaxdamagepercent": {17, 0},
	"itemmindamagepercent": {18, 0}, "enhanceddamage": {18, 0},
	"tohit":                {19},
	"toblock":              {20},
	"plusmindamage":        {21, 1},
	"mindamage":            {21},
	"plusmaxdamage":        {22, 1},
	"maxdamage":            {22},
	"secondarymindamage":   {23},
	"secondarymaxdamage":   {24},
	"damagepercent":        {25},
	"manarecovery":         {26},
	"manarecoverybonus":    {27},
	"staminarecoverybonus": {28},
	"lastexp":              {29},
	"nextexp":              {30},

	"armorclass": {31}, "defense": {31},
	"plusdefense": {31, 0},

	"armorclassvsmissile":   {32},
	"armorclassvshth":       {33},
	"normaldamagereduction": {34},
	"magicdamagereduction":  {35},
	"damageresist":          {36},
	"magicresist":           {37},
	"maxmagicresist":        {38},
	"fireresist":            {39},
	"maxfireresist":         {40},
	"lightresist":           {41},
	"maxlightresist":        {42},
	"coldresist":            {43},
	"maxcoldresist":         {44},
	"poisonresist":          {45},
	"maxpoisonresist":       {46},
	"damageaura":            {47},
	"firemindam":            {48},
	"firemaxdam":            {49},
	"lightmindam":           {50},
	"lightmaxdam":           {51},
	"magicmindam":           {52},
	"magicmaxdam":           {53},
	"coldmindam":            {54},
	"coldmaxdam":            {55},
	"coldlength":            {56},
	"poisondamage":          {57, 1},
	"poisonmindam":          {57},
	"poisonmaxdam":          {58},
	"poisonlength":          {59},
	"lifedrainmindam":       {60}, "lifeleech": {60},
	"lifedrainmaxdam": {61},
	"manadrainmindam": {62}, "manaleech": {62},
	"manadrainmaxdam":          {63},
	"stamdrainmindam":          {64},
	"stamdrainmaxdam":          {65},
	"stunlength":               {66},
	"velocitypercent":          {67},
	"attackrate":               {68},
	"otheranimrate":            {69},
	"quantity":                 {70},
	"value":                    {71},
	"durability":               {72},
	"maxdurability":            {73},
	"hpregen":                  {74},
	"itemmaxdurabilitypercent": {75},
	"itemmaxhppercent":         {76},
	"itemmaxmanapercent":       {77},
	"itemattackertakesdamage":  {78},
	"itemgoldbonus":            {79},
	"itemmagicbonus":           {80},
	"itemknockback":            {81},
	"itemtimeduration":         {82},

	"itemaddclassskills":  {83},
	"itemaddamazonskills": {83, 0}, "amazonskills": {83, 0},
	"itemaddsorceressskills": {83, 1}, "sorceressskills": {83, 1},
	"itemaddnecromancerskills": {83, 2}, "necromancerskills": {83, 2},
	"itemaddpaladinskills": {83, 3}, "paladinskills": {83, 3},
	"itemaddbarbarianskills": {83, 4}, "barbarianskills": {83, 4},
	"itemadddruidskills": {83, 5}, "druidskills": {83, 5},
	"itemaddassassinskills": {83, 6}, "assassinskills": {83, 6},

	"unsentparam1":           {84},
	"itemaddexperience":      {85},
	"itemhealafterkill":      {86},
	"itemreducedprices":      {87},
	"itemdoubleherbduration": {88},
	"itemlightradius":        {89},
	"itemlightcolor":         {90},
	"itemreqpercent":         {91},
	"itemlevelreq":           {92},
	"itemfasterattackrate":   {93}, "ias": {93},
	"itemlevelreqpct":        {94},
	"lastblockframe":         {95},
	"itemfastermovevelocity": {96}, "frw": {96},

	// oskill
	"itemnonclassskill": {97},
	// Amazon
	"plusskillcriticalstrike": {97, 9},
	"plusskillguidedarrow":    {97, 22},
	// Sorceress
	"plusskillteleport": {97, 54},
	// Barbarian
	"plusskillbattleorders":  {97, 149},
	"plusskillbattlecommand": {97, 155},
	"plusskillbattlecry":     {97, 146},
	// Druid
	"plusskillwerewolf":      {97, 223},
	"plusskillshapeshifting": {97, 224}, "plusskilllycanthropy": {97, 224},
	"plusskillsummonspiritwolf": {97, 227},
	"plusskillferalrage":        {97, 232},

	"state":                {98},
	"itemfastergethitrate": {99}, "fhr": {99},
	"monsterplayercount":        {100},
	"skillpoisonoverridelength": {101},
	"itemfasterblockrate":       {102}, "fbr": {102},
	"skillbypassundead":  {103},
	"skillbypassdemons":  {104},
	"itemfastercastrate": {105}, "fcr": {105},
	"skillbypassbeasts": {106},

	"itemsingleskill": {107},
	// Amazon skills
	"skillmagicarrow":      {107, 6},
	"skillfirearrow":       {107, 7},
	"skillinnersight":      {107, 8},
	"skillcriticalstrike":  {107, 9},
	"skilljab":             {107, 10},
	"skillcoldarrow":       {107, 11},
	"skillmultipleshot":    {107, 12},
	"skilldodge":           {107, 13},
	"skillpowerstrike":     {107, 14},
	"skillpoisonjavelin":   {107, 15},
	"skillexplodingarrow":  {107, 16},
	"skillslowmissiles":    {107, 17},
	"skillavoid":           {107, 18},
	"skillimpale":          {107, 19},
	"skilllightningbolt":   {107, 20},
	"skillicearrow":        {107, 21},
	"skillguidedarrow":     {107, 22},
	"skillpenetrate":       {107, 23},
	"skillchargedstrike":   {107, 24},
	"skillplaguejavelin":   {107, 25},
	"skillstrafe":          {107, 26},
	"skillimmolationarrow": {107, 27},
	"skilldecoy":           {107, 28},
	"skillevade":           {107, 29},
	"skillfend":            {107, 30},
	"skillfreezingarrow":   {107, 31},
	"skillvalkyrie":        {107, 32},
	"skillpierce":          {107, 33},
	"skilllightningstrike": {107, 34},
	"skilllightningfury":   {107, 35},
	// Sorceress skills
	"skillfirebolt":         {107, 36},
	"skillwarmth":           {107, 37},
	"skillchargedbolt":      {107, 38},
	"skillicebolt":          {107, 39},
	"skillfrozenarmor":      {107, 40},
	"skillinferno":          {107, 41},
	"skillstaticfield":      {107, 42},
	"skilltelekinesis":      {107, 43},
	"skillfrostnova":        {107, 44},
	"skilliceblast":         {107, 45},
	"skillblaze":            {107, 46},
	"skillfireball":         {107, 47},
	"skillnova":             {107, 48},
	"skilllightning":        {107, 49},
	"skillshiverarmor":      {107, 50},
	"skillfirewall":         {107, 51},
	"skillenchant":          {107, 52},
	"skillchainlightning":   {107, 53},
	"skillteleport":         {107, 54},
	"skillglacialspike":     {107, 55},
	"skillmeteor":           {107, 56},
	"skillthunderstorm":     {107, 57},
	"skillenergyshield":     {107, 58},
	"skillblizzard":         {107, 59},
	"skillchillingarmor":    {107, 60},
	"skillfiremastery":      {107, 61},
	"skillhydra":            {107, 62},
	"skilllightningmastery": {107, 63},
	"skillfrozenorb":        {107, 64},
	"skillcoldmastery":      {107, 65},
	// Necromancer skills
	"skillamplifydamage":   {107, 66},
	"skillteeth":           {107, 67},
	"skillbonearmor":       {107, 68},
	"skillskeletonmastery": {107, 69},
	"skillraiseskeleton":   {107, 70},
	"skilldimvision":       {107, 71},
	"skillweaken":          {107, 72},
	"skillpoisondagger":    {107, 73},
	"skillcorpseexplosion": {107, 74},
	"skillclaygolem":       {107, 75},
	"skillironmaiden":      {107, 76},
	"skillterror":          {107, 77},
	"skillbonewall":        {107, 78},
	"skillgolemmastery":    {107, 79},
	"skillskeletalmage":    {107, 80},
	"skillconfuse":         {107, 81},
	"skilllifetap":         {107, 82},
	"skillpoisonexplosion": {107, 83},
	"skillbonespear":       {107, 84},
	"skillbloodgolem":      {107, 85},
	"skillattract":         {107, 86},
	"skilldecrepify":       {107, 87},
	"skillboneprison":      {107, 88},
	"skillsummonresist":    {107, 89},
	"skillirongolem":       {107, 90},
	"skilllowerresist":     {107, 91},
	"skillpoisonnova":      {107, 92},
	"skillbonespirit":      {107, 93},
	"skillfiregolem":       {107, 94},
	"skillrevive":          {107, 95},
	// Paladin skills
	"skillsacrifice":        {107, 96},
	"skillsmite":            {107, 97},
	"skillmight":            {107, 98},
	"skillprayer":           {107, 99},
	"skillresistfire":       {107, 100},
	"skillholybolt":         {107, 101},
	"skillholyfire":         {107, 102},
	"skillthorns":           {107, 103},
	"skilldefiance":         {107, 104},
	"skillresistcold":       {107, 105},
	"skillzeal":             {107, 106},
	"skillcharge":           {107, 107},
	"skillblessedaim":       {107, 108},
	"skillcleansing":        {107, 109},
	"skillresistlightning":  {107, 110},
	"skillvengeance":        {107, 111},
	"skillblessedhammer":    {107, 112},
	"skillconcentration":    {107, 113},
	"skillholyfreeze":       {107, 114},
	"skillvigor":            {107, 115},
	"skillconversion":       {107, 116},
	"skillholyshield":       {107, 117},
	"skillholyshock":        {107, 118},
	"skillsanctuary":        {107, 119},
	"skillmeditation":       {107, 120},
	"skillfistoftheheavens": {107, 121},
	"skillfanaticism":       {107, 122},
	"skillconviction":       {107, 123},
	"skillredemption":       {107, 124},
	"skillsalvation":        {107, 125},
	// Barbarian skills
	"skillbash":              {107, 126},
	"skillswordmastery":      {107, 127},
	"skillaxemastery":        {107, 128},
	"skillmacemastery":       {107, 129},
	"skillhowl":              {107, 130},
	"skillfindpotion":        {107, 131},
	"skillleap":              {107, 132},
	"skilldoubleswing":       {107, 133},
	"skillpolearmmastery":    {107, 134},
	"skillthrowingmastery":   {107, 135},
	"skillspearmastery":      {107, 136},
	"skilltaunt":             {107, 137},
	"skillshout":             {107, 138},
	"skillstun":              {107, 139},
	"skilldoublethrow":       {107, 140},
	"skillincreasedstamina":  {107, 141},
	"skillfinditem":          {107, 142},
	"skillleapattack":        {107, 143},
	"skillconcentrate":       {107, 144},
	"skillironskin":          {107, 145},
	"skillbattlecry":         {107, 146},
	"skillfrenzy":            {107, 147},
	"skillincreasedspeed":    {107, 148},
	"skillbattleorders":      {107, 149},
	"skillgrimward":          {107, 150},
	"skillwhirlwind":         {107, 151},
	"skillberserk":           {107, 152},
	"skillnaturalresistance": {107, 153},
	"skillwarcry":            {107, 154},
	"skillbattlecommand":     {107, 155},
	// Druid skills
	"skillraven":            {107, 221},
	"skillpoisoncreeper":    {107, 222},
	"skillwerewolf":         {107, 223},
	"skilllycanthropy":      {107, 224},
	"skillfirestorm":        {107, 225},
	"skilloaksage":          {107, 226},
	"skillsummonspiritwolf": {107, 227},
	"skillwerebear":         {107, 228},
	"skillmoltenboulder":    {107, 229},
	"skillarcticblast":      {107, 230},
	"skillcarrionvine":      {107, 231},
	"skillferalrage":        {107, 232},
	"skillmaul":             {107, 233},
	"skillfissure":          {107, 234},
	"skillcyclonearmor":     {107, 235},
	"skillheartofwolverine": {107, 236},
	"skillsummondirewolf":   {107, 237},
	"skillrabies":           {107, 238},
	"skillfireclaws":        {107, 239},
	"skilltwister":          {107, 240},
	"skillsolarcreeper":     {107, 241},
	"skillhunger":           {107, 242},
	"skillshockwave":        {107, 243},
	"skillvolcano":          {107, 244},
	"skilltornado":          {107, 245},
	"skillspiritofbarbs":    {107, 246},
	"skillsummongrizzly":    {107, 247},
	"skillfury":             {107, 248},
	"skillarmageddon":       {107, 249},
	"skillhurricane":        {107, 250},
	// Assassin skills
	"skillfireblast":         {107, 251},
	"skillclawmastery":       {107, 252},
	"skillpsychichammer":     {107, 253},
	"skilltigerstrike":       {107, 254},
	"skilldragontalon":       {107, 255},
	"skillshockweb":          {107, 256},
	"skillbladesentinel":     {107, 257},
	"skillburstofspeed":      {107, 258},
	"skillfistsoffire":       {107, 259},
	"skilldragonclaw":        {107, 260},
	"skillchargedboltsentry": {107, 261},
	"skillwakeoffire":        {107, 262},
	"skillweaponblock":       {107, 263},
	"skillcloakofshadows":    {107, 264},
	"skillcobrastrike":       {107, 265},
	"skillbladefury":         {107, 266},
	"skillfade":              {107, 267},
	"skillshadowwarrior":     {107, 268},
	"skillclawsofthunder":    {107, 269},
	"skilldragontail":        {107, 270},
	"skilllightningsentry":   {107, 271},
	"skillwakeofinferno":     {107, 272},
	"skillmindblast":         {107, 273},
	"skillbladesofice":       {107, 274},
	"skilldragonflight":      {107, 275},
	"skilldeathsentry":       {107, 276},
	"skillbladeshield":       {107, 277},
	"skillvenom":             {107, 278},
	"skillshadowmaster":      {107, 279},
	"skillphoenixstrike":     {107, 280},

	"itemrestinpeace":              {108},
	"curseresistance":              {109},
	"itempoisonlengthresist":       {110},
	"itemnormaldamage":             {111},
	"itemhowl":                     {112},
	"itemstupidity":                {113},
	"itemdamagetomana":             {114},
	"itemignoretargetac":           {115},
	"itemfractionaltargetac":       {116},
	"itempreventheal":              {117},
	"itemhalffreezeduration":       {118},
	"itemtohitpercent":             {119},
	"itemdamagetargetac":           {120},
	"itemdemondamagepercent":       {121},
	"itemundeaddamagepercent":      {122},
	"itemdemontohit":               {123},
	"itemundeadtohit":              {124},
	"itemthrowable":                {125},
	"itemelemskill":                {126},
	"itemallskills":                {127},
	"itemattackertakeslightdamage": {128},
	"ironmaidenlevel":              {129},
	"lifetaplevel":                 {130},
	"thornspercent":                {131},
	"bonearmor":                    {132},
	"bonearmormax":                 {133},
	"itemfreeze":                   {134},
	"itemopenwounds":               {135},
	"itemcrushingblow":             {136},
	"itemkickdamage":               {137},
	"itemmanaafterkill":            {138},
	"itemhealafterdemonkill":       {139},
	"itemextrablood":               {140},
	"itemdeadlystrike":             {141},
	"itemabsorbfirepercent":        {142},
	"itemabsorbfire":               {143},
	"itemabsorblightpercent":       {144},
	"itemabsorblight":              {145},
	"itemabsorbmagicpercent":       {146},
	"itemabsorbmagic":              {147},
	"itemabsorbcoldpercent":        {148},
	"itemabsorbcold":               {149},
	"itemslow":                     {150},

	"itemaura":          {151},
	"mightaura":         {151, 98},
	"holyfireaura":      {151, 102},
	"thornsaura":        {151, 103},
	"defianceaura":      {151, 104},
	"concentrationaura": {151, 113},
	"holyfreezeaura":    {151, 114},
	"vigoraura":         {151, 115},
	"holyshockaura":     {151, 118},
	"sanctuaryaura":     {151, 119},
	"meditationaura":    {151, 120},
	"fanaticismaura":    {151, 122},
	"convictionaura":    {151, 123},
	"redemptionaura":    {151, 124},

	"itemindestructible":             {152},
	"itemcannotbefrozen":             {153},
	"itemstaminadrainpct":            {154},
	"itemreanimate":                  {155},
	"itempierce":                     {156},
	"itemmagicarrow":                 {157},
	"itemexplosivearrow":             {158},
	"itemthrowmindamage":             {159},
	"itemthrowmaxdamage":             {160},
	"itemskillhandofathena":          {161},
	"itemskillstaminapercent":        {162},
	"itemskillpassivestaminapercent": {163},
	"itemskillconcentration":         {164},
	"itemskillenchant":               {165},
	"itemskillpierce":                {166},
	"itemskillconviction":            {167},
	"itemskillchillingarmor":         {168},
	"itemskillfrenzy":                {169},
	"itemskilldecrepify":             {170},
	"itemskillarmorpercent":          {171},
	"alignment":                      {172},
	"target0":                        {173},
	"target1":                        {174},
	"goldlost":                       {175},
	"conversionlevel":                {176},
	"conversionmaxhp":                {177},
	"unitdooverlay":                  {178},
	"attackvsmontype":                {179},
	"damagevsmontype":                {180},
	"fade":                           {181},
	"armoroverridepercent":           {182},
	"unused183":                      {183},
	"unused184":                      {184},
	"unused185":                      {185},
	"unused186":                      {186},
	"itempiercecoldimmunity":         {187},

	"itemaddskilltab":               {188},
	"itemaddbowandcrossbowskilltab": {188, 0}, "bowandcrossbowskilltab": {188, 0},
	"itemaddpassiveandmagicskilltab": {188, 1}, "passiveandmagicskilltab": {188, 1},
	"itemaddjavelinandspearskilltab": {188, 2}, "javelinandspearskilltab": {188, 2},
	"itemaddfireskilltab": {188, 8}, "fireskilltab": {188, 8},
	"itemaddlightningskilltab": {188, 9}, "lightningskilltab": {188, 9},
	"itemaddcoldskilltab": {188, 10}, "coldskilltab": {188, 10},
	"itemaddcursesskilltab": {188, 16}, "cursesskilltab": {188, 16},
	"itemaddpoisonandboneskilltab": {188, 17}, "poisonandboneskilltab": {188, 17},
	"itemaddnecromancersummoningskilltab": {188, 18}, "necromancersummoningskilltab": {188, 18},
	"itemaddpalicombatskilltab": {188, 24}, "palicombatskilltab": {188, 24},
	"itemaddoffensiveaurasskilltab": {188, 25}, "offensiveaurasskilltab": {188, 25},
	"itemadddefensiveaurasskilltab": {188, 26}, "defensiveaurasskilltab": {188, 26},
	"itemaddbarbcombatskilltab": {188, 32}, "barbcombatskilltab": {188, 32},
	"itemaddmasteriesskilltab": {188, 33}, "masteriesskilltab": {188, 33},
	"itemaddwarcriesskilltab": {188, 34}, "warcriesskilltab": {188, 34},
	"itemadddruidsummoningskilltab": {188, 40}, "druidsummoningskilltab": {188, 40},
	"itemaddshapeshiftingskilltab": {188, 41}, "shapeshiftingskilltab": {188, 41},
	"itemaddelementalskilltab": {188, 42}, "elementalskilltab": {188, 42},
	"itemaddtrapsskilltab": {188, 48}, "trapsskilltab": {188, 48},
	"itemaddshadowdisciplinesskilltab": {188, 49}, "shadowdisciplinesskilltab": {188, 49},
	"itemaddmartialartsskilltab": {188, 50}, "martialartsskilltab": {188, 50},

	"itempiercefireimmunity":   {189},
	"itempiercelightimmunity":  {190},
	"itempiercepoisonimmunity": {191},
	"itempiercedamageimmunity": {192},
	"itempiercemagicimmunity":  {193},
	"itemnumsockets":           {194}, "sockets": {194},
	"itemskillonattack":      {195, 1},
	"itemskillonattacklevel": {195, 2},
	"itemskillonkill":        {196, 1},
	"itemskillonkilllevel":   {196, 2},
	"itemskillondeath":       {197, 1},
	"itemskillondeathlevel":  {197, 2},

	"itemskillonhit":      {198, 1},
	"itemskillonhitlevel": {198, 2},
	"amplifydamageonhit":  {198, 4225},

	"itemskillonlevelup":      {199, 1},
	"itemskillonleveluplevel": {199, 2},
	"unused200":               {200},
	"itemskillongethit":       {201, 1},
	"itemskillongethitlevel":  {201, 2},
	"unused202":               {202},
	"unused203":               {203},

	"itemchargedskill":      {204, 1},
	"itemchargedskilllevel": {204, 2},
	"teleportcharges":       {204, 3461},

	"unused204":                    {205},
	"unused205":                    {206},
	"unused206":                    {207},
	"unused207":                    {208},
	"unused208":                    {209},
	"unused209":                    {210},
	"unused210":                    {211},
	"unused211":                    {212},
	"unused212":                    {213},
	"itemarmorperlevel":            {214},
	"itemarmorpercentperlevel":     {215},
	"itemhpperlevel":               {216},
	"itemmanaperlevel":             {217},
	"itemmaxdamageperlevel":        {218},
	"itemmaxdamagepercentperlevel": {219},
	"itemstrengthperlevel":         {220},
	"itemdexterityperlevel":        {221},
	"itemenergyperlevel":           {222},
	"itemvitalityperlevel":         {223},
	"itemtohitperlevel":            {224},
	"itemtohitpercentperlevel":     {225},
	"itemcolddamagemaxperlevel":    {226},
	"itemfiredamagemaxperlevel":    {227},
	"itemltngdamagemaxperlevel":    {228},
	"itempoisdamagemaxperlevel":    {229},
	"itemresistcoldperlevel":       {230},
	"itemresistfireperlevel":       {231},
	"itemresistltngperlevel":       {232},
	"itemresistpoisperlevel":       {233},
	"itemabsorbcoldperlevel":       {234},
	"itemabsorbfireperlevel":       {235},
	"itemabsorbltngperlevel":       {236},
	"itemabsorbpoisperlevel":       {237},
	"itemthornsperlevel":           {238},
	"itemfindgoldperlevel":         {239},
	"itemfindmagicperlevel":        {240},
	"itemregenstaminaperlevel":     {241},
	"itemstaminaperlevel":          {242},
	"itemdamagedemonperlevel":      {243},
	"itemdamageundeadperlevel":     {244},
	"itemtohitdemonperlevel":       {245},
	"itemtohitundeadperlevel":      {246},
	"itemcrushingblowperlevel":     {247},
	"itemopenwoundsperlevel":       {248},
	"itemkickdamageperlevel":       {249},
	"itemdeadlystrikeperlevel":     {250},
	"itemfindgemsperlevel":         {251},
	"itemreplenishdurability":      {252},
	"itemreplenishquantity":        {253},
	"itemextrastack":               {254},
	"itemfinditem":                 {255},
	"itemslashdamage":              {256},
	"itemslashdamagepercent":       {257},
	"itemcrushdamage":              {258},
	"itemcrushdamagepercent":       {259},
	"itemthrustdamage":             {260},
	"itemthrustdamagepercent":      {261},
	"itemabsorbslash":              {262},
	"itemabsorbcrush":              {263},
	"itemabsorbthrust":             {264},
	"itemabsorbslashpercent":       {265},
	"itemabsorbcrushpercent":       {266},
	"itemabsorbthrustpercent":      {267},
	"itemarmorbytime":              {268},
	"itemarmorpercentbytime":       {269},
	"itemhpbytime":                 {270},
	"itemmanabytime":               {271},
	"itemmaxdamagebytime":          {272},
	"itemmaxdamagepercentbytime":   {273},
	"itemstrengthbytime":           {274},
	"itemdexteritybytime":          {275},
	"itemenergybytime":             {276},
	"itemvitalitybytime":           {277},
	"itemtohitbytime":              {278},
	"itemtohitpercentbytime":       {279},
	"itemcolddamagemaxbytime":      {280},
	"itemfiredamagemaxbytime":      {281},
	"itemltngdamagemaxbytime":      {282},
	"itempoisdamagemaxbytime":      {283},
	"itemresistcoldbytime":         {284},
	"itemresistfirebytime":         {285},
	"itemresistltngbytime":         {286},
	"itemresistpoisbytime":         {287},
	"itemabsorbcoldbytime":         {288},
	"itemabsorbfirebytime":         {289},
	"itemabsorbltngbytime":         {290},
	"itemabsorbpoisbytime":         {291},
	"itemfindgoldbytime":           {292},
	"itemfindmagicbytime":          {293},
	"itemregenstaminabytime":       {294},
	"itemstaminabytime":            {295},
	"itemdamagedemonbytime":        {296},
	"itemdamageundeadbytime":       {297},
	"itemtohitdemonbytime":         {298},
	"itemtohitundeadbytime":        {299},
	"itemcrushingblowbytime":       {300},
	"itemopenwoundsbytime":         {301},
	"itemkickdamagebytime":         {302},
	"itemdeadlystrikebytime":       {303},
	"itemfindgemsbytime":           {304},
	"itempiercecold":               {305},
	"itempiercefire":               {306},
	"itempierceltng":               {307},
	"itempiercepois":               {308},
	"itemdamagevsmonster":          {309},
	"itemdamagepercentvsmonster":   {310},
	"itemtohitvsmonster":           {311},
	"itemtohitpercentvsmonster":    {312},
	"itemacvsmonster":              {313},
	"itemacpercentvsmonster":       {314},
	"firelength":                   {315},
	"burningmin":                   {316},
	"burningmax":                   {317},
	"progressivedamage":            {318},
	"progressivesteal":             {319},
	"progressiveother":             {320},
	"progressivefire":              {321},
	"progressivecold":              {322},
	"progressivelightning":         {323},
	"itemextracharges":             {324},
	"progressivetohit":             {325},
	"poisoncount":                  {326},
	"damageframerate":              {327},
	"pierceidx":                    {328},
	"passivefiremastery":           {329},
	"passiveltngmastery":           {330},
	"passivecoldmastery":           {331},
	"passivepoismastery":           {332},
	"passivefirepierce":            {333},
	"passiveltngpierce":            {334},
	"passivecoldpierce":            {335},
	"passivepoispierce":            {336},
	"passivecriticalstrike":        {337},
	"passivedodge":                 {338},
	"passiveavoid":                 {339},
	"passiveevade":                 {340},
	"passivewarmth":                {341},
	"passivemasterymeleeth":        {342},
	"passivemasterymeleedmg":       {343},
	"passivemasterymeleecrit":      {344},
	"passivemasterythrowth":        {345},
	"passivemasterythrowdmg":       {346},
	"passivemasterythrowcrit":      {347},
	"passiveweaponblock":           {348},
	"passivesummonresist":          {349},
	"modifierlistskill":            {350},
	"modifierlistlevel":            {351},
	"lastsenthppct":                {352},
	"sourceunittype":               {353},
	"sourceunitid":                 {354},
	"shortparam1":                  {355},
	"questitemdifficulty":          {356},
	"passivemagmastery":            {357},
	"passivemagpierce":             {358},
	"skillcooldown":                {359},
	"skillmissiledamagescale":      {360},

	// Doesnt really exists, but is calculated in getStatEx
	"allres": {555},
}
