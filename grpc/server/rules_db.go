package server

var dndRulesDB = []struct {
	keywords []string
	rule     string
}{
	// ── CORE MECHANICS ──────────────────────────────────────────────────────────
	{
		keywords: []string{"attack", "hit", "roll", "attack roll", "to hit"},
		rule:     "ATTACK ROLLS: Roll d20 + ability modifier + proficiency bonus (if proficient). Melee attacks use STR (or DEX for finesse weapons). Ranged attacks use DEX. A natural 20 is a Critical Hit (double damage dice). A natural 1 always misses. Two-Weapon Fighting lets you use a Bonus Action to attack with your off-hand light weapon (no modifier to damage unless negative).",
	},
	{
		keywords: []string{"advantage", "disadvantage", "roll twice"},
		rule:     "ADVANTAGE & DISADVANTAGE: Roll 2d20 and take the higher (Advantage) or lower (Disadvantage) result. Multiple sources of Advantage don't stack — you always roll exactly two dice. If you have both Advantage and Disadvantage simultaneously, they cancel out regardless of how many sources each has.",
	},
	{
		keywords: []string{"saving throw", "save", "resistance", "immunity", "vulnerability"},
		rule:     "SAVING THROWS & DAMAGE TYPES: Roll d20 + ability modifier + proficiency bonus (if proficient). DCs: Trivial (5), Easy (10), Medium (15), Hard (20), Very Hard (25), Nearly Impossible (30). Resistance halves damage; Immunity negates it; Vulnerability doubles it. Damage types: Acid, Bludgeoning, Cold, Fire, Force, Lightning, Necrotic, Piercing, Poison, Psychic, Radiant, Slashing, Thunder.",
	},
	{
		keywords: []string{"ability score", "modifier", "attribute", "str", "dex", "con", "int", "wis", "cha"},
		rule:     "ABILITY SCORES & MODIFIERS: Scores range 1–30; modifier = (score − 10) / 2, rounded down. STR: Athletics, melee attacks, carrying capacity. DEX: Acrobatics, Stealth, Sleight of Hand, ranged attacks, AC in light/no armor. CON: HP per level, concentration checks. INT: Arcana, History, Investigation, Nature, Religion. WIS: Animal Handling, Insight, Medicine, Perception, Survival. CHA: Deception, Intimidation, Performance, Persuasion.",
	},
	{
		keywords: []string{"proficiency", "proficiency bonus", "level", "experience", "xp", "leveling"},
		rule:     "PROFICIENCY BONUS & LEVELING: Proficiency bonus by tier — Levels 1–4: +2 (Tier 1, local heroes); Levels 5–8: +3 (Tier 2, heroes of the realm); Levels 9–12: +4 (Tier 3, masters of the realm); Levels 13–16: +5 (Tier 4, masters of the world); Levels 17–20: +6 (Tier 4 apex). XP thresholds double roughly each level. Milestone leveling is a popular alternative where the DM awards levels at story beats.",
	},
	{
		keywords: []string{"initiative", "combat", "turn", "round", "action", "bonus action", "reaction", "free action"},
		rule:     "COMBAT STRUCTURE: Initiative = d20 + DEX modifier; ties broken by DEX score, then coin flip. Each turn: Move (up to Speed), Action (Attack, Cast Spell, Dash, Disengage, Dodge, Help, Hide, Ready, Search, Use Object), optional Bonus Action, optional Free Interaction. Reactions (e.g., Opportunity Attack, Shield spell) occur outside your turn. One round = 6 seconds of in-game time.",
	},
	{
		keywords: []string{"opportunity attack", "leaving", "movement", "disengage"},
		rule:     "OPPORTUNITY ATTACKS: Triggered when a hostile creature you can see moves out of your reach without Disengaging. Uses your Reaction. Disengage action prevents triggering opportunity attacks for the rest of your turn. Forced movement (shoves, spells) does not trigger opportunity attacks.",
	},
	{
		keywords: []string{"grapple", "grappling", "shove", "shoving", "restrained", "prone"},
		rule:     "GRAPPLING & SHOVING: Both use your Attack action (replace one attack). Grapple: contested Athletics vs. Athletics or Acrobatics — success = target gains Grappled condition (speed 0). Shove: contested Athletics vs. Athletics or Acrobatics — choose to knock Prone or push 5 ft. Grappled creatures can use their action to escape (same contest). Grappler's speed is not reduced by Grapple itself, but they drag the target at half speed.",
	},
	{
		keywords: []string{"cover", "half cover", "three-quarters cover", "full cover", "hiding", "invisible"},
		rule:     "COVER & CONCEALMENT: Half cover (+2 AC and DEX saves) — low walls, creatures. Three-quarters cover (+5 AC and DEX saves) — arrow slits, thick tree trunks. Total cover — can't be directly targeted. Lightly Obscured (dim light, fog): Disadvantage on Perception. Heavily Obscured (darkness, dense fog): effectively Blinded. Invisible attackers have Advantage; attacks against them have Disadvantage.",
	},
	{
		keywords: []string{"death", "dying", "unconscious", "death saving throw", "stabilize", "massive damage", "instant death"},
		rule:     "DEATH & DYING: Drop to 0 HP = unconscious + start death saves. Each turn: d20 — 10+ success, 9− failure. 3 successes = stable; 3 failures = dead. Nat 1 = 2 failures; Nat 20 = regain 1 HP. Any damage while at 0 HP adds 1 failure; a Critical Hit adds 2. Instant death if damage exceeds HP maximum as negative. Healer's Kit (action) stabilizes without a check.",
	},
	{
		keywords: []string{"skill", "check", "ability check", "passive", "passive perception", "investigation", "perception"},
		rule:     "SKILL CHECKS & PASSIVE SCORES: Roll d20 + modifier + proficiency (if proficient). Expertise doubles proficiency bonus. Passive score = 10 + modifiers (used for background checks without rolling). Passive Perception is the most common — the DM uses it to determine if characters notice hidden threats. Contested checks: both sides roll, higher wins (ties go to the active participant).",
	},
	{
		keywords: []string{"rest", "short rest", "long rest", "heal", "recovery", "hit dice", "exhaustion"},
		rule:     "RESTING: Short Rest (1+ hour, no strenuous activity): Spend Hit Dice to recover HP (roll HD + CON modifier per die). Long Rest (8+ hours, max 2 of light activity): Recover all HP, regain half max Hit Dice (min 1), regain all spell slots, reset death saves. One long rest per 24 hours. Exhaustion is removed by one level per long rest. Interrupted long rests require starting over.",
	},
	{
		keywords: []string{"condition", "blinded", "charmed", "deafened", "exhausted", "frightened", "incapacitated", "invisible", "paralyzed", "petrified", "poisoned", "prone", "restrained", "stunned", "unconscious"},
		rule:     "CONDITIONS REFERENCE: Blinded: auto-fail sight checks, attack Disadvantage, attacks against have Advantage. Charmed: can't attack charmer, charmer has Advantage on social checks. Frightened: Disadvantage on checks/attacks while source visible, can't move closer. Incapacitated: no actions or reactions. Paralyzed: Incapacitated + auto-fail STR/DEX saves, attacks have Advantage, crits within 5 ft. Petrified: weight ×10, Incapacitated, Resistant to all damage, immune to poison/disease. Prone: Disadvantage on attacks; melee attacks against have Advantage, ranged have Disadvantage; stand up costs half movement. Restrained: speed 0, attack Disadvantage, DEX saves Disadvantage, attacks against have Advantage. Stunned: Incapacitated + can't move + auto-fail STR/DEX saves. Unconscious: Incapacitated + prone + auto-fail STR/DEX saves + attacks have Advantage + crits within 5 ft.",
	},
	{
		keywords: []string{"exhaustion", "fatigue", "tired"},
		rule:     "EXHAUSTION LEVELS: 1 — Disadvantage on ability checks. 2 — Speed halved. 3 — Disadvantage on attack rolls and saving throws. 4 — HP maximum halved. 5 — Speed reduced to 0. 6 — Death. Each long rest removes one level. Exhaustion sources: forced march, extreme heat/cold, starvation, certain spells and monsters.",
	},
	{
		keywords: []string{"concentration", "concentrating", "lose concentration"},
		rule:     "CONCENTRATION: Only one concentration spell active at a time — casting another ends the first. When taking damage while concentrating: CON save DC = max(10, half damage taken). Distraction (being shoved, rough seas) may also require a DC 10 CON save. War Caster feat grants Advantage on concentration saves. Incapacitated or dying immediately breaks concentration.",
	},
	// ── SPELLCASTING ────────────────────────────────────────────────────────────
	{
		keywords: []string{"spell", "casting", "spellcasting", "magic", "spell slot", "cantrip", "ritual"},
		rule:     "SPELLCASTING FUNDAMENTALS: Spell save DC = 8 + proficiency + spellcasting modifier. Spell attack bonus = proficiency + spellcasting modifier. Upcasting: expend a higher slot to enhance certain spells. Cantrips require no slot and scale with character level (not class level) at levels 5, 11, 17. Ritual casting (takes 10 extra minutes) requires the Ritual Caster feat or a class feature — no slot expended.",
	},
	{
		keywords: []string{"spell component", "verbal", "somatic", "material", "component pouch", "arcane focus"},
		rule:     "SPELL COMPONENTS: V (Verbal) — must speak; silenced creatures can't cast. S (Somatic) — free hand required; shields/two-handed weapons may block. M (Material) — specific items consumed if listed with cost; a Component Pouch or Arcane Focus replaces non-costly materials. Subtle Spell (Sorcerer Metamagic) removes V and S requirements.",
	},
	{
		keywords: []string{"wild magic", "sorcerer", "sorcery points", "metamagic", "twinned", "quickened", "subtle"},
		rule:     "SORCERER: Uses Charisma. Sorcery Points (SP) = level. Flexible Casting: 1 SP per spell slot level (1st = 1SP, 2nd = 3SP, 3rd = 5SP, 4th = 6SP, 5th = 7SP). Metamagic options include: Careful (save others from area spells), Distant (double range), Empowered (reroll damage dice), Extended (double duration), Heightened (Disadvantage on first save), Quickened (cast as Bonus Action), Subtle (no V/S), Twinned (target second creature).",
	},
	{
		keywords: []string{"warlock", "eldritch blast", "pact magic", "eldritch invocation", "patron", "pact boon"},
		rule:     "WARLOCK: Uses Charisma. Pact Magic: few spell slots (max 5th level) but all regained on Short Rest. Eldritch Invocations customize abilities — e.g., Agonizing Blast (+CHA to Eldritch Blast), Devil's Sight (see in magical darkness 120 ft), Misty Visions (silent image at will). Pact Boons: Blade (summoned weapon), Chain (improved familiar), Tome (three cantrips + rituals). Mystic Arcanum grants one spell per tier (6th/7th/8th/9th) once per long rest.",
	},
	{
		keywords: []string{"wizard", "arcane recovery", "spellbook", "spell scroll", "prepared spells"},
		rule:     "WIZARD: Uses Intelligence. Knows spells via Spellbook (starts with 6 + INT modifier; copy more for 50gp + 2hr each). Prepares INT modifier + level spells each long rest. Arcane Recovery (short rest, once/day): regain spell slots totaling half wizard level (rounded up), no slot above 5th. Spell Mastery (18th): cast one 1st and one 2nd level spell at will. Signature Spells (20th): two 3rd-level spells as if always prepared, once per short/long rest without slot.",
	},
	{
		keywords: []string{"cleric", "divine domain", "channel divinity", "turn undead", "divine intervention"},
		rule:     "CLERIC: Uses Wisdom. Prepares WIS modifier + level spells. Channel Divinity (regained on short rest at 6th level, long rest before): Turn Undead (Wisdom save or Turned for 1 minute), plus Domain ability. Destroy Undead at 5th level (auto-destroy low-CR undead). Divine Intervention (20th): 100% chance to call on deity. Domains grant bonus spells always prepared and armor/weapon proficiencies vary by domain.",
	},
	{
		keywords: []string{"druid", "wild shape", "beast", "moon druid", "animal form", "nature"},
		rule:     "DRUID: Uses Wisdom. Wild Shape (twice per short rest): transform into beasts with CR ≤ level/4 (max CR 1 at lvl 4, CR 1/2 at lvl 2); can't speak, cast spells, or use equipment. Moon Druid (Circle of the Moon) can transform into CR ≤ level/3 and combat forms starting level 2. Retain personality/mental scores; gain beast's physical scores and HP. Revert when HP drops to 0 (excess damage carries over), unconscious, or willingly.",
	},
	{
		keywords: []string{"paladin", "lay on hands", "divine smite", "aura", "oath", "sacred weapon"},
		rule:     "PALADIN: Uses Charisma. Lay on Hands (HP pool = level × 5): heal HP or cure disease/poison (5 HP from pool). Divine Smite: expend a spell slot after hitting — deal 2d8 radiant per slot level (min 1st = 2d8, +1d8 per level above 1st, +1d8 vs undead/fiends; max 5d8). Auras (10th: 30 ft): Aura of Protection (+CHA to all saving throws for allies). Sacred Weapon, Wrathful Smite, etc. vary by Oath.",
	},
	{
		keywords: []string{"ranger", "hunter", "favored enemy", "natural explorer", "beast companion", "hunter's mark"},
		rule:     "RANGER: Uses Wisdom for spells (half-caster). Favored Enemy: Advantage on Survival checks to track, Intelligence checks to recall info. Natural Explorer: double proficiency on INT/WIS checks in favored terrain, not slowed by difficult terrain. Hunter's Mark (concentration): mark target, deal +1d6 damage on hits, Bonus Action to switch. Beast Master: bond with an animal companion that follows commands.",
	},
	{
		keywords: []string{"rogue", "sneak attack", "cunning action", "uncanny dodge", "evasion", "thieves tools", "lockpicking"},
		rule:     "ROGUE: Uses DEX primarily. Sneak Attack: once per turn, +1d6 per 2 rogue levels, requires Advantage OR an ally within 5 ft of target (and no Disadvantage). Cunning Action (Bonus Action): Dash, Disengage, or Hide. Uncanny Dodge (5th): halve damage from one attacker per round as Reaction. Evasion (7th): DEX save for half damage instead takes none (fail takes half). Expertise: double proficiency on two chosen skills.",
	},
	{
		keywords: []string{"fighter", "action surge", "second wind", "extra attack", "fighting style", "indomitable"},
		rule:     "FIGHTER: Second Wind (Bonus Action, short/long rest): regain 1d10 + fighter level HP. Action Surge (short/long rest): gain one additional Action this turn (2 uses at 17th). Extra Attack: 2 attacks at 5th, 3 at 11th, 4 at 20th. Fighting Styles: Archery (+2 ranged), Defense (+1 AC in armor), Dueling (+2 melee with one hand free), Great Weapon (+reroll 1s and 2s on damage), Protection (impose Disadvantage on attacker with shield as Reaction), Two-Weapon. Indomitable (9th): reroll a failed saving throw.",
	},
	{
		keywords: []string{"barbarian", "rage", "reckless attack", "unarmored defense", "brutal critical", "bear totem"},
		rule:     "BARBARIAN: Rage (Bonus Action; uses = 2+level or unlimited at 20th): +2 melee damage (scales), Resistance to bludgeoning/piercing/slashing, Advantage on STR checks/saves; ends if you don't attack or take damage by end of turn, or willingly. Reckless Attack: Advantage on melee attacks this turn, enemies have Advantage against you until next turn. Unarmored Defense: AC = 10 + DEX + CON (no armor). Brutal Critical (9th): one extra damage die on crits.",
	},
	{
		keywords: []string{"bard", "bardic inspiration", "jack of all trades", "song of rest", "magical secrets"},
		rule:     "BARD: Uses Charisma. Bardic Inspiration (Bonus Action; uses = CHA modifier, regain on long rest or short rest at 5th): grant d6 (scales to d12 at 15th) to ally's roll within 60 ft, used within 10 minutes. Jack of All Trades (2nd): add half proficiency to non-proficient ability checks. Song of Rest (2nd): allies regain extra HD worth of HP on short rests. Magical Secrets (10th): learn 2 spells from any class list.",
	},
	{
		keywords: []string{"monk", "ki", "ki point", "flurry of blows", "patient defense", "step of the wind", "stunning strike", "unarmored movement"},
		rule:     "MONK: Uses WIS and DEX. Ki Points = level (regain on short rest). Flurry of Blows (1 ki): two Bonus Action unarmed strikes. Patient Defense (1 ki): Dodge as Bonus Action. Step of the Wind (1 ki): Disengage or Dash as Bonus Action, double jump distance. Stunning Strike (1 ki): after hitting, target makes CON save or Stunned until end of your next turn. Martial Arts: use DEX for monk weapons/unarmed; Bonus Action unarmed strike. Unarmored Defense: AC = 10 + DEX + WIS.",
	},
	// ── RACES ───────────────────────────────────────────────────────────────────
	{
		keywords: []string{"race", "species", "human", "elf", "dwarf", "halfling", "dragonborn", "gnome", "half-elf", "half-orc", "tiefling", "aasimar", "tabaxi", "tortle", "genasi"},
		rule:     "RACES OVERVIEW: Human: +1 to all scores, one extra skill proficiency, one feat (Variant Human). Elf: +2 DEX, Darkvision 60 ft, Trance (4-hr rest), Fey Ancestry (Advantage vs charm, immune to sleep magic). Dwarf: +2 CON, Darkvision 60 ft, Dwarven Resilience (Advantage vs poison, resistance to poison damage), Stonecunning. Halfling: +2 DEX, Lucky (reroll 1s on d20 rolls), Brave (Advantage vs Frightened), Halfling Nimbleness (move through larger creature's space). Dragonborn: +2 STR +1 CHA, Breath Weapon (based on draconic ancestry, 2d6 scaling damage, DEX/CON save), Damage Resistance.",
	},
	{
		keywords: []string{"half-elf", "half orc", "tiefling", "gnome", "aasimar"},
		rule:     "MORE RACES: Half-Elf: +2 CHA +1 to two others, Darkvision, Fey Ancestry, two skill proficiencies. Half-Orc: +2 STR +1 CON, Darkvision, Relentless Endurance (drop to 1 HP instead of 0, once per long rest), Savage Attacks (+1 die on melee crits). Tiefling: +2 CHA +1 INT, Darkvision, Hellish Resistance (fire), Infernal Legacy (Thaumaturgy cantrip; Hellish Rebuke 1/day at 3rd; Darkness 1/day at 5th). Gnome: +2 INT, Darkvision 60 ft, Gnome Cunning (Advantage on INT/WIS/CHA saves vs magic). Aasimar: +2 CHA, Darkvision, Celestial Resistance (necrotic and radiant), Healing Hands, Light Bearer.",
	},
	// ── MONSTERS & ENCOUNTER DESIGN ─────────────────────────────────────────────
	{
		keywords: []string{"monster", "creature", "beast", "undead", "fiend", "celestial", "dragon", "giant", "humanoid", "aberration", "construct", "elemental", "fey", "monstrosity", "ooze", "plant", "swarm"},
		rule:     "CREATURE TYPES: Aberration (alien intellect — mind flayers, beholders), Beast (natural animals), Celestial (good-aligned extraplanar — angels, unicorns), Construct (created — golems, animated armor), Dragon (true dragons and kin — wyverns, pseudodragons), Elemental (planar — elementals, mephits), Fey (Feywild origin — hags, dryads), Fiend (evil extraplanar — demons, devils, yugoloths), Giant (large humanoids — ogres, trolls, giants), Humanoid (most civilized races), Monstrosity (magical oddities — griffons, manticores), Ooze (formless — gelatinous cube, black pudding), Plant (animated flora — shambling mound, treant), Undead (animated dead — skeletons, vampires, liches).",
	},
	{
		keywords: []string{"challenge rating", "cr", "xp", "encounter", "deadly", "easy", "medium", "hard", "encounter budget"},
		rule:     "CHALLENGE RATING & ENCOUNTER BUDGETS: CR represents the challenge for a party of 4 at that level. XP by CR: CR 0 = 10 XP, CR 1/8 = 25, CR 1/4 = 50, CR 1/2 = 100, CR 1 = 200, CR 2 = 450, CR 5 = 1800, CR 10 = 5900, CR 15 = 13000, CR 20 = 25000. Encounter difficulty thresholds per character per day — Easy/Medium/Hard/Deadly. Multiply total XP by modifier based on number of monsters: 1 (×1), 2 (×1.5), 3-6 (×2), 7-10 (×2.5), 11-14 (×3), 15+ (×4).",
	},
	{
		keywords: []string{"legendary", "legendary action", "legendary resistance", "lair action", "lair", "regional effect", "boss"},
		rule:     "LEGENDARY CREATURES: Legendary Resistance (3/day): choose to succeed on a failed saving throw. Legendary Actions (usually 3): taken at end of other creatures' turns, reset at start of their turn — typically Move, Attack, or a special ability costing 1-3 actions. Lair Actions (initiative count 20, if in lair): environmental effects that never target the same creature twice in a row. Regional Effects: the creature's presence warps the surrounding region (e.g., dragon's lair causes minor earthquakes, plant overgrowth).",
	},
	{
		keywords: []string{"dragon", "breath weapon", "dragon age", "wyrmling", "ancient dragon", "adult dragon", "chromatic", "metallic"},
		rule:     "DRAGONS: Age categories — Wyrmling, Young, Adult, Ancient (increasing CR, size, and abilities). Chromatic (evil): Black (acid), Blue (lightning), Green (poison), Red (fire), White (cold). Metallic (good): Brass/Copper (fire/acid), Bronze/Silver (lightning/cold), Gold (fire). Breath Weapon: recharges on 5-6, uses a cone or line area (DEX or CON save). Dragons have Legendary Resistance, Frightful Presence (Wisdom save or Frightened 1 min, 120 ft), and Legendary Actions. Ancient dragons have lair and regional effects.",
	},
	{
		keywords: []string{"undead", "vampire", "zombie", "skeleton", "lich", "ghoul", "specter", "wraith", "ghost", "wight", "revenant"},
		rule:     "UNDEAD KEY TRAITS: Most undead are immune to poison and psychic damage, exhaustion, being poisoned, frightened, and don't need food/water/air. Vampires: Regeneration (20 HP/turn if not in sunlight/running water), Charm (Wisdom save), Vampire Weaknesses (sunlight 20 radiant/turn, running water 20 acid/turn, repelled by garlic/holy symbols). Liches: Legendary creature, Phylactery (reform 1d10 days after death), paralyzing touch. Ghouls: Paralyzing Claws (CON save DC 10 or Paralyzed 1 min, doesn't affect elves).",
	},
	{
		keywords: []string{"demon", "devil", "fiend", "hell", "abyss", "summoning", "planar binding"},
		rule:     "FIENDS — DEMONS VS DEVILS: Demons (Chaotic Evil, The Abyss): immune to lightning, poison, fire; resistant to cold, and non-magical physical. Devils (Lawful Evil, Nine Hells): immune to fire and poison; resistant to cold and non-magical physical. Both have Darkvision 120 ft and telepathy. Summoned fiends return to their plane when reduced to 0 HP. Planar Binding (spell) can compel a summoned fiend to serve for the spell's duration.",
	},
	{
		keywords: []string{"giant", "ogre", "troll", "troll regeneration", "hill giant", "stone giant", "fire giant", "frost giant", "cloud giant", "storm giant"},
		rule:     "GIANTS & TROLLS: Giants have a social hierarchy (Ordning) from lowest to highest: Hill, Stone, Frost, Fire, Cloud, Storm. Trolls regenerate 10 HP per turn (stopped by acid or fire damage that round). Trolls grow back severed limbs unless burned. Giants typically have rock throwing (long range), great strength, and resistances based on type. Fire Giants are immune to fire; Frost Giants are immune to cold.",
	},
	// ── EQUIPMENT & ITEMS ───────────────────────────────────────────────────────
	{
		keywords: []string{"armor", "armor class", "ac", "light armor", "medium armor", "heavy armor", "shield", "stealth disadvantage"},
		rule:     "ARMOR & AC: Unarmored: 10 + DEX modifier (or class features). Light Armor (DEX no cap): Padded (AC 11), Leather (AC 11), Studded Leather (AC 12). Medium Armor (max +2 DEX): Hide (AC 12), Chain Shirt (AC 13), Scale Mail (AC 14), Breastplate (AC 14), Half Plate (AC 15). Heavy Armor (no DEX): Ring Mail (AC 14), Chain Mail (AC 16, STR 13), Splint (AC 17, STR 15), Plate (AC 18, STR 15). Shield: +2 AC (cannot use two-handed weapon). Stealth Disadvantage: Padded, Scale Mail, Splint, Chain Mail, Plate, Ring Mail.",
	},
	{
		keywords: []string{"weapon", "sword", "dagger", "bow", "crossbow", "axe", "finesse", "thrown", "reach", "two-handed", "versatile"},
		rule:     "WEAPON PROPERTIES: Finesse (use STR or DEX), Thrown (ranged option with listed range), Two-Handed (requires both hands), Versatile (one-handed or two-handed damage), Reach (+5 ft range), Light (off-hand fighting), Loading (one shot per Action regardless of attacks), Heavy (Disadvantage for Small/Tiny creatures). Common weapons: Dagger 1d4 (finesse, thrown 20/60), Shortsword 1d6 (finesse), Longsword 1d8/1d10 versatile, Greatsword 2d6 (two-handed), Longbow 1d8 (range 150/600, heavy, two-handed).",
	},
	{
		keywords: []string{"magic item", "attunement", "attuned", "magic weapon", "magic armor", "rarity", "uncommon", "rare", "very rare", "legendary item"},
		rule:     "MAGIC ITEMS & ATTUNEMENT: Rarities: Common, Uncommon, Rare, Very Rare, Legendary, Artifact. Attunement: some items require a Short Rest to attune; you can be attuned to max 3 items. Removing attunement requires another Short Rest. Magic weapons and armor overcome non-magical resistance/immunity to their damage type. Cursed items can't be removed without Remove Curse. Identifying items: short rest of study or Identify spell.",
	},
	{
		keywords: []string{"potion", "healing potion", "potion of healing", "antitoxin", "poison", "venom"},
		rule:     "POTIONS & CONSUMABLES: Potion of Healing: 2d4+2 HP (Action to drink; Bonus Action if given to you by another). Greater Healing: 4d4+4. Superior: 8d4+8. Supreme: 10d4+20. Antitoxin: Advantage on CON saves vs poison for 1 hour. Poisons: Injury (applied to weapon, DC CON save on hit), Ingested (must consume, delayed effect), Inhaled (affects all in area), Contact (touch to skin). Poisoned condition: Disadvantage on attack rolls and ability checks.",
	},
	// ── EXPLORATION & ENVIRONMENT ────────────────────────────────────────────────
	{
		keywords: []string{"movement", "speed", "difficult terrain", "climbing", "swimming", "flying", "jumping", "crawling"},
		rule:     "MOVEMENT & TERRAIN: Difficult terrain costs double movement per foot. Climbing costs double movement (triple if difficult). Swimming costs double (check CON in rough water). Flying creatures knocked prone fall unless they can hover. Long jump (running): STR score in feet (halved from standing). High jump (running): 3 + STR modifier in feet. Crawling: double movement cost. You can split movement before and after actions.",
	},
	{
		keywords: []string{"falling", "fall damage", "fall", "falling damage"},
		rule:     "FALLING: Take 1d6 bludgeoning per 10 feet fallen (max 20d6 / 200 ft). Land prone unless you avoid taking damage (e.g., Feather Fall). Falling into water from great height: DM may allow DEX save to reduce damage. Xanathar's Guide variant: if fall lasts until end of turn, you hit the ground at the start of your next turn (relevant for reactions that can interrupt falls).",
	},
	{
		keywords: []string{"suffocation", "drowning", "holding breath", "underwater", "water combat"},
		rule:     "SUFFOCATION & UNDERWATER COMBAT: Hold breath for 1 + CON modifier minutes (min 30 sec). After that: survival for CON modifier rounds (min 1) before dropping to 0 HP. Underwater: slashing and bludgeoning ranged weapon attacks have Disadvantage; fire-based spells may not function. Creatures without swim speed must swim (Athletics) and have Disadvantage on melee weapon attacks unless they have a swim speed or are a water-breathing creature.",
	},
	{
		keywords: []string{"light", "darkness", "darkvision", "blindsight", "tremorsense", "truesight", "dim light", "bright light"},
		rule:     "VISION & LIGHT: Bright Light — normal. Dim Light (lightly obscured): Disadvantage on Perception. Darkness (heavily obscured): effectively Blinded. Darkvision: see in darkness as dim light, dim as bright — no color. Blindsight: perceive without sight within range (ignores invisibility, darkness). Tremorsense: detect vibrations through ground contact. Truesight: see through illusions, invisibility, magical darkness, true form of shapeshifters, into Ethereal Plane.",
	},
	{
		keywords: []string{"trap", "traps", "disarm trap", "detect trap", "thieves tools", "search"},
		rule:     "TRAPS: Spotting: Passive Perception or active Investigation/Perception check against Trap DC. Disarming: Thieves' Tools (DEX) or specified method. Trigger types: Pressure plate, tripwire, proximity, magical sensor. Common effects: Pit (falling damage), Arrow trap (ranged attack vs AC), Poison gas (CON save), Magical trap (various — Glyph of Warding). Failing by 5+ on a disarm check triggers the trap. Some magical traps require Dispel Magic.",
	},
	{
		keywords: []string{"stealth", "hiding", "hidden", "sneak", "hide action"},
		rule:     "STEALTH & HIDING: Hide Action: make a DEX (Stealth) check; result = DC for Perception checks against you. Must be out of clear sight — lightly obscured counts. Staying hidden: revealed if you attack, cast a non-subtle spell, or a creature succeeds on Perception against your Stealth. Invisible creatures still make Stealth checks to remain undetected by hearing. Passive Perception is used for unaware observers.",
	},
	{
		keywords: []string{"mounted combat", "mount", "horse", "warhorse", "flying mount"},
		rule:     "MOUNTED COMBAT: Mount must be Large or larger and willing. Mount acts on your Initiative. Controlled mount: limited to Dash, Disengage, Dodge. Independent mount: acts on its own Initiative. If mount is knocked prone, rider dismounts in its space (DC 10 DEX save or also falls prone). If mount is moved against its will, Rider saves DC 10 DEX or dismounts prone. Rider can choose to fall off to reduce damage (Acrobatics check).",
	},
	// ── SOCIAL & ROLEPLAY ───────────────────────────────────────────────────────
	{
		keywords: []string{"social", "persuasion", "deception", "intimidation", "charisma check", "npc attitude", "friendly", "hostile", "indifferent", "charm"},
		rule:     "SOCIAL INTERACTION: NPC attitudes — Hostile (may attack or refuse), Indifferent (unconcerned), Friendly (willing to help). Persuasion (CHA): honest appeals to improve attitude. Deception (CHA): false information or impersonation. Intimidation (CHA or STR): threats and coercion. Insight (WIS): detect lies or read motives. Good roleplay can grant Advantage. Bribery may require no check if the offer is sufficient. Failed checks may lock out further social attempts.",
	},
	{
		keywords: []string{"alignment", "lawful", "chaotic", "good", "evil", "neutral", "true neutral"},
		rule:     "ALIGNMENT: Two axes — Law/Chaos and Good/Evil. Lawful Good: honor and compassion (paladin archetype). Neutral Good: does good without rigid structure. Chaotic Good: free spirit guided by conscience. Lawful Neutral: order above all. True Neutral: balance. Chaotic Neutral: follows whims. Lawful Evil: methodical villainy. Neutral Evil: serves self without scruple. Chaotic Evil: wanton destruction. Alignment guides roleplay but doesn't restrict class choices (5e is more flexible than prior editions).",
	},
	{
		keywords: []string{"language", "common", "elvish", "dwarvish", "draconic", "infernal", "abyssal", "celestial", "thieves cant", "druidic"},
		rule:     "LANGUAGES: Standard: Common, Dwarvish, Elvish, Giant, Gnomish, Goblin, Halfling, Orc. Exotic: Abyssal (demons), Celestial (celestials), Draconic (dragons), Deep Speech (aberrations), Infernal (devils), Primordial (elementals — Aquan, Auran, Ignan, Terran), Sylvan (fey), Undercommon (Underdark traders). Secret: Thieves' Cant (rogues — encoded messages in normal speech), Druidic (druids only). INT modifier grants extra languages at character creation from Background.",
	},
	{
		keywords: []string{"background", "feature", "backstory", "bond", "flaw", "ideal", "personality"},
		rule:     "BACKGROUNDS: Each grants two skill proficiencies, a tool or language proficiency, starting equipment, and a unique Feature (narrative benefit — e.g., Soldier's Military Rank grants deference from soldiers; Criminal's Criminal Contact provides underworld information). Also define Personality Traits (2), Ideal (1), Bond (1), Flaw (1). These guide roleplay and can be used for Inspiration awards by the DM.",
	},
	{
		keywords: []string{"inspiration", "inspired", "heroic inspiration"},
		rule:     "INSPIRATION: The DM awards Inspiration for great roleplay, honoring your character's traits/ideals/bonds/flaws, or exceptional creativity. You either have it or you don't (can't stack). Spend it: gain Advantage on one attack roll, ability check, or saving throw. You can give your Inspiration to another player. Some optional rules grant Inspiration for using spell components or honoring alignment.",
	},
	// ── MULTICLASSING & FEATS ───────────────────────────────────────────────────
	{
		keywords: []string{"multiclass", "multiclassing", "dip", "prestige", "dual class"},
		rule:     "MULTICLASSING: Requires minimum ability scores in both current and new class (e.g., Paladin needs STR 13 + CHA 13). Each new class level grants that class's features except starting equipment. Proficiencies gained are limited (usually armor/weapon subset). Spell slots combined from a table based on total spellcaster levels (half-casters like Paladin/Ranger count as half level). Spells prepared still use each class's formula separately. Extra Attack doesn't stack between Fighter, Paladin, Ranger.",
	},
	{
		keywords: []string{"feat", "feats", "ability score improvement", "asi", "war caster", "lucky", "sentinel", "sharpshooter", "great weapon master", "polearm master"},
		rule:     "FEATS (replacing ASI at certain levels): Lucky: 3 luck points/day — roll extra d20 for attack/check/save, choose which to use (even after seeing result). Sentinel: opportunity attacks reduce speed to 0; can make opportunity attack when enemy uses Disengage; can attack when ally is attacked within reach. War Caster: Advantage on concentration saves; can perform somatic components while holding weapons/shield; cast spells as opportunity attacks. Great Weapon Master: -5 attack/+10 damage (Heavy weapons); Bonus Action attack on crit or killing blow. Sharpshooter: same -5/+10 for ranged; ignore half/three-quarters cover; no Disadvantage at long range.",
	},
	// ── PLANES & COSMOLOGY ───────────────────────────────────────────────────────
	{
		keywords: []string{"plane", "outer plane", "inner plane", "astral", "ethereal", "feywild", "shadowfell", "nine hells", "abyss", "mount celestia", "mechanus"},
		rule:     "PLANES OF EXISTENCE: Material Plane (prime world). Feywild (echo of beauty and magic — fey creatures, timelessness). Shadowfell (echo of death and despair — undead, Shadar-kai). Ethereal Plane (overlaps Material, accessed by Etherealness). Astral Plane (silvery void between planes — githyanki, psychic damage danger). Inner Planes: Fire, Water, Air, Earth, Positive/Negative Energy. Outer Planes (aligned to alignment axis): Mount Celestia (LG), Mechanus (LN), Nine Hells/Baator (LE), The Abyss (CE), Limbo (CN), Elysium (NG), Ysgard (CG).",
	},
	// ── MISCELLANEOUS ────────────────────────────────────────────────────────────
	{
		keywords: []string{"encumbrance", "carrying capacity", "weight", "push", "drag", "lift"},
		rule:     "ENCUMBRANCE: Carrying Capacity = STR score × 15 lbs. Push/Drag/Lift = STR score × 30 lbs (speed drops to 5 ft above carrying capacity). Variant Encumbrance: Encumbered (> STR×5 lbs) = −10 speed; Heavily Encumbered (> STR×10 lbs) = −20 speed + Disadvantage on DEX/STR/CON ability checks, attacks, and saving throws. Most tables ignore this rule for simplicity.",
	},
	{
		keywords: []string{"food", "water", "starvation", "foraging", "survival", "ration"},
		rule:     "FOOD, WATER & SURVIVAL: Food: 1 lb/day. Miss a day: CON save (DC 10, +1 per prior failure) or gain exhaustion level. Water: 1 gallon/day (2 in hot weather). Half ration: CON save DC 15 or exhaustion. Foraging: WIS (Survival) vs DC 10 (forest), DC 20 (desert), fail = nothing found that day. Rations last 1 day and automatically succeed food requirements.",
	},
	{
		keywords: []string{"downtime", "crafting", "carousing", "training", "spellbook copying", "running a business"},
		rule:     "DOWNTIME ACTIVITIES: Crafting: 5 gp materials per day toward item cost (must be proficient). Practicing a Profession: maintain lifestyle without costs. Recuperating: 3+ days, CON save DC 15 to remove a lingering injury or disease. Research: find information (cost + Investigation check). Training a Language/Tool: 250 days + 1 gp/day. Carousing: 1d3 days, 1d6×10 gp — social connections or trouble (1d8 table). Spellbook copying: 2 hours + 50 gp per spell level.",
	},
	{
		keywords: []string{"disease", "curse", "lycanthropy", "werewolf", "vampire spawn", "lycanthrope"},
		rule:     "DISEASE, CURSES & LYCANTHROPY: Diseases are contracted by various means; Lesser Restoration cures most mundane diseases. Curses require Remove Curse (3rd-level spell). Lycanthropy: contracted when a lycanthrope bites and reduces you to 0 HP; CON save DC 8 + lycanthrope's proficiency + CON modifier or contract the curse. At each full moon the afflicted transforms uncontrollably. Greater Restoration or Remove Curse within 3 days can cure it. Vampire Spawn: humanoid reduced to 0 HP by a vampire and not raised as a different undead rises as a Vampire Spawn under the vampire's control.",
	},
}
