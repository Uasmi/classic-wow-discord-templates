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

Bonus: Rotten Ghoul

## Abilities explained
###### Disclaimer: though we really wanted to create lot's of cool abilities, there are some limitations in Discord that we can't surpass (or possibly don't know how). If you have a suggestion just create and Issue on this GitHub and we'll take a look if it's possible
#### ![](https://vignette.wikia.nocookie.net/wowwiki/images/9/91/Spell_shadow_cripple.png/revision/latest?cb=20060930190414) Banish
*-banish @nickname* to send the target to the other realm! (text channel...) Other channels are hidden. While target is banished, it is impossible to Dispel, Polymorph or Bubble the target.
#### ![](https://vignette.wikia.nocookie.net/wowwiki/images/9/92/Spell_shaman_improvedreincarnation.png/revision/latest?cb=20100901165909) Reincarnation
Shaman is able to Reincarnate after control spells was used on him. This will trigger the cooldown of the caster, but will not affect the Shaman himself (works on Polymorph, Freezing Trap, Banish)
#### Stealth
*-stealth* to dissapear and *-unstealth* to appear again! 
###### P.S. Those filthy casters can't get you but be careful of freezing traps!
#### Dispel
The most selfless people recieve ability to remove control spells. Feel free to help your friend, and use *-dispel @nickname* to remove control spells.
###### Don't tell me you're a shadow priest, we all know this doesn't work.
#### Bubble
Do I need to explain this? Use *-bubble @nickname* to put a Bubble on ~~someone~~ yourself
#### Polymorph
*-polymorph @nickname*

We all like to turn people into sheeps... Baah :sheep:
#### Freezing Trap
*-trap* to set a freezing trap. 
The next person who'll type in chat will get into Freezing Trap. You'll recieve a DM that the trap was set. 
###### Tss, don't tell anyone...
#### Teldrassil
Separate text channel for Druids
###### Come on, you have too many cool things in the game, separate Teldrassil for you is not that bad!
#### Taunt
This one is a surprise. To trigger the command just type *-taunt* and see what will happen.
###### Make sure you joined the voice channel

## Requirements and prework:
You'll need to do the following steps to make this work:
1. Setup a discord server
2. Add YAGPDB to your server: https://yagpdb.xyz
3. Make custom roles such as:
//add pic here

## Setting up:
Here will be steps to add commands to your server
There are several 

Login to YAGPDB portal (through discord authorization) and choose your server. Then on the left you'll see Core dropdown, press it and then choose custom commands. After doing that you need to copy the code from *.go* files at this repository (CTRL-C/CTRL-V will work fine) and paste it to the portal one by one. 
The only important file is a trapRegexTracker, which is using Regex. You need to change the trigger to regex, and set command to .*
Here's and example:
