# WoW Classic - Discord custom commands
Hello and welcome to this repo, which contains custom commands for your discord server, which you can freely implement and have fun!

## Classes and Interactions:
Adding this commands to your Discord server will add class specific features! Each person on your server will be able to choose a class and use it's ability:

1. ![](https://vignette.wikia.nocookie.net/wowwiki/images/9/94/Warlock_Icon.gif/revision/latest?cb=20070911030126)Warlock - Banish
2. ![](https://vignette.wikia.nocookie.net/wowwiki/images/9/98/Shaman_Icon.gif/revision/latest?cb=20070911030053)Shaman - Reincarnation 
3. ![](https://vignette.wikia.nocookie.net/wowwiki/images/3/38/Rogue_Icon.gif/revision/latest?cb=20070911030015)Rogue - Stealth
4. ![](https://vignette.wikia.nocookie.net/wowwiki/images/1/17/Priest_Icon.gif/revision/latest?cb=20070911025947)Priest - Dispel
5. ![](https://vignette.wikia.nocookie.net/wowwiki/images/a/a5/Paladin_Icon.gif/revision/latest?cb=20070911025906)Paladin - Bubble
6. ![](https://vignette.wikia.nocookie.net/wowwiki/images/0/07/Mage_Icon.gif/revision/latest?cb=20070911025832)Mage - Polymorph
7. ![](https://vignette.wikia.nocookie.net/wowwiki/images/b/b6/Hunter_Icon.gif/revision/latest?cb=20070911025740)Hunter - Freezing Trap
8. ![](https://vignette.wikia.nocookie.net/wowwiki/images/6/6b/Druid_Icon.gif/revision/latest?cb=20070911025603)Druid - Teldrassil
9. ![](https://vignette.wikia.nocookie.net/wowwiki/images/b/bc/Warrior_Icon.gif/revision/latest?cb=20070911030206)Warrior - Taunt

###### Bonus: Rotten Ghoul - if someone is acting bad, promote him to Rotten Ghoul and send him to Eastern Plaguelands

## Abilities explained
###### Disclaimer: though we really wanted to create lot's of cool abilities, there are some limitations in Discord that we can't surpass (or possibly don't know how). If you have a suggestion just create and Issue on this Repo and we'll take a look if it's possible
#### ![](https://vignette.wikia.nocookie.net/wowwiki/images/9/91/Spell_shadow_cripple.png/revision/latest?cb=20060930190414)  Banish
*-banish @nickname* to send the target to the other realm! (text channel...) Other channels are hidden. While target is banished, it is impossible to Dispel, Polymorph or Bubble the target.
#### ![](https://vignette.wikia.nocookie.net/wowwiki/images/9/92/Spell_shaman_improvedreincarnation.png/revision/latest?cb=20100901165909)  Reincarnation
Shaman is able to Reincarnate after control spells was used on him. This will trigger the cooldown of the caster, but will not affect the Shaman himself (works on Polymorph, Freezing Trap, Banish)
#### ![](https://vignette.wikia.nocookie.net/wowwiki/images/f/f8/Ability_stealth.png/revision/latest?cb=20051006101655) Stealth
*-stealth* to dissapear and *-unstealth* to appear again! 
###### P.S. Those filthy casters can't get you but be careful of freezing traps!
#### ![](https://vignette.wikia.nocookie.net/wowwiki/images/5/5e/Spell_holy_dispelmagic.png/revision/latest?cb=20060930054627) Purify
The most selfless people recieve ability to remove control spells. Feel free to help your friend, and use *-purify @nickname* to remove control spells.
###### Don't tell me you're a shadow priest, we all know this doesn't work.
#### ![](https://vignette.wikia.nocookie.net/wowwiki/images/1/18/Spell_holy_divineshield.png/revision/latest?cb=20111101153930) Bubble
Do I need to explain this? Use *-bubble @nickname* to put a Bubble on ~~someone~~ yourself
#### ![](https://wow.zamimg.com/images/wow/icons/large/spell_nature_polymorph.jpg) Polymorph
*-polymorph @nickname*

We all like to turn people into sheeps... Baah :sheep:
#### ![](https://wow.zamimg.com/images/wow/icons/large/spell_frost_chainsofice.jpg) Freezing Trap
*-trap* to set a freezing trap. 
The next person who'll type in chat will get into Freezing Trap. You'll recieve a DM that the trap was set. 
###### Tss, don't tell anyone...
#### <img src="https://wow.zamimg.com/uploads/screenshots/normal/114418-teldrassil-teldrassil-concept-art-not-a-stump.jpg" width="80" height="45" /> Teldrassil
Separate text channel for Druids
###### Come on, Druids. You have too many cool things in the game! Separate Teldrassil is not that bad!
#### ![](https://wow.zamimg.com/images/wow/icons/large/spell_nature_reincarnation.jpg) Taunt
This one is a surprise. To trigger the command just type *-taunt* and see what will happen.
###### Make sure you joined the voice channel!

## Requirements and prework:
Now let's jump into setting things up.

You'll need to do the following steps to make this work:
#### 1. Setup a discord server: https://discordapp.com/
#### 2. Add YAGPDB to your server: https://yagpdb.xyz
#### 3. Make custom roles such as below:
<img src="https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/properRoles.jpg?raw=true" width="400" height="450" />

You can get hex code colours for classes here: https://wow.gamepedia.com/Class_colors

Set *Sheep, Frozen and Banished* roles to read-only in settings of the role
Set *Stealth* role color to ```#2C2F33``` so it will look like this: <img src="https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/stealth.jpg?raw=true" width="100" height="25" />

#### 4. Setup a separate channel where you can let people choose a Class (self-assignable roles):
Follow this cool youtube video to do that:
https://www.youtube.com/watch?v=84nndebMEGo
## Setting up:
Now you'll have to add commands to your server.
Make sure you selected your server at Yagpdb portal (you can see it on top right corner)
To do that, jump to yagpdb portal, and in *Core* drop-down choose *Custom Commands*:

![](https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/cc.jpg?raw=true)

## Adding Go files
Start to add commands by selecting each .go file from this repository, and copying code to yagpdb portal (Press on *.go* file, Copy the code and Paste to custom command). Don't forget to set the command to *Require at least one of the roles in the following lists* and assign appropriate Class (ex., banish = Warlock)

Here's and example:

![](https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/addingCC.gif?raw=true)

Repeat the procces for (keep the command triggers same as file names):
1. banish - Warlock
2. bubble - Paladin
3. polymorph - Mage
4. purify - Priest
5. rolemanager - keep role fields as **None**
6. stealth - Rogue
7. trap - Hunter
8. unstealth - Rogue
9. taunt - Warrior
10. regexTrapTracker - **this one is important, see below**

The only important file is a trapRegexTracker, which is using Regex. You need to change the **Trigger Type** to regex, and set **Trigger** to .*: 

![](https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/regex.jpg?raw=true)

Create a separate text channel for druids called Teldrassil:

![](https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/teldrassil.jpg?raw=true)

## Adding sounds
Sounds are also included, you can get them at sounds folder. Download them and then Upload to Yagpdb portal.
Go to *Fun* dropdown and choose *Soundboard*:

![](https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/sb.jpg?raw=true)

**Please keep the names for sounds as below:**

![](https://github.com/Uasmi/classic-wow-discord-templates/blob/master/pics/sounds.jpg?raw=true)

## Changing Cooldown and Ability effect length:
You can see two variables named 
```$cooldown```
```$abilityTime``` at each Command.
You can set preffered duration in seconds for each ability. By default, the abilites have:
1. Cooldown 300s, Ability Effect Length 30s: Bubble, Banish
2. Cooldown 300s, Ability Effect Length 15s: Trap, Polymorph
2. Cooldown 300s: Purify
2. Cooldown 30s: Taunt

Stealth doesn't have a cooldown. 


## Have fun!
That's it!
Feel free to ask questions at Issues. :)
