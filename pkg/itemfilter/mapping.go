package itemfilter

import (
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/nip"
)

var qualities = map[string]item.Quality{
	nip.QualityLowQuality: item.QualityLowQuality,
	nip.QualityNormal:     item.QualityNormal,
	nip.QualitySuperior:   item.QualitySuperior,
	nip.QualityMagic:      item.QualityMagic,
	nip.QualitySet:        item.QualitySet,
	nip.QualityRare:       item.QualityRare,
	nip.QualityUnique:     item.QualityUnique,
	nip.QualityCrafted:    item.QualityCrafted,
}

// Source: https://github.com/blizzhackers/kolbot/blob/mainline/d2bs/kolbot/libs/NTItemAlias.dbl
var stats = map[string][]int{
	"itemhealafterdemonkill": {139},
	"itemextrablood":         {140},
	"itemdeadlystrike":       {141},
	"itemabsorbfirepercent":  {142},
	"itemabsorbfire":         {143},
	"itemabsorblightpercent": {144},
	"itemabsorblight":        {145},
	"itemabsorbmagicpercent": {146},
	"itemabsorbmagic":        {147},
	"itemabsorbcoldpercent":  {148},
	"itemabsorbcold":         {149},
	"itemslow":               {150},

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
	"unused187":                      {187},

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

	"unused189":      {189},
	"unused190":      {190},
	"unused191":      {191},
	"unused192":      {192},
	"unused193":      {193},
	"itemnumsockets": {194}, "sockets": {194},
	"itemskillonattack": {195},
	"itemskillonkill":   {196},
	"itemskillondeath":  {197},

	"itemskillonhit":     {198},
	"amplifydamageonhit": {198, 4225},

	"itemskillonlevelup": {199},
	"unused200":          {200},
	"itemskillongethit":  {201},
	"unused202":          {202},
	"unused203":          {203},

	"itemchargedskill": {204},
	"teleportcharges":  {204, 3461},

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
}
