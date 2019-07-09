{{$usertag:= userArg (index .CmdArgs 0)}}
{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
  Oops, you're on cooldown. 
{{else}}
  You were polymorphed {{$usertag.Username}}!
  :sheep:Baaah:sheep: 
  {{giveRoleName $usertag "Sheep"}}
  {{takeRoleName $usertag "Sheep" (30)}}
  {{dbSetExpire .User.ID ("usedCC") ("True") (100)}}
  
  {{if targetHasRoleName $usertag "Mage"}}
    {{takeRoleName $usertag "Mage"}}
    {{sleep 30}}
    {{giveRolename $usertag "Mage"}}
  {{else if targetHasRoleName $usertag "Warrior"}}
    {{takeRoleName $usertag "Warrior"}}
    {{sleep 30}}
    {{giveRolename $usertag "Warrior"}}
  {{else if targetHasRoleName $usertag "Warlock"}}
    {{takeRoleName $usertag "Warlock"}}
    {{sleep 30}}
    {{giveRolename $usertag "Warlock"}}
  {{else if targetHasRoleName $usertag "Shaman"}}
    {{takeRoleName $usertag "Shaman"}}
    {{sleep 30}}
    {{giveRolename $usertag "Shaman"}}
  {{else if targetHasRoleName $usertag "Rogue"}}
    {{takeRoleName $usertag "Rogue"}}
    {{sleep 30}}
    {{giveRolename $usertag "Rogue"}}
  {{else if targetHasRoleName $usertag "Priest"}}
    {{takeRoleName $usertag "Priest"}}
    {{sleep 30}}
    {{giveRolename $usertag "Priset"}}
  {{else if targetHasRoleName $usertag "Paladin"}}
    {{takeRoleName $usertag "Paladin"}}
    {{sleep 30}}
    {{giveRolename $usertag "Paladin"}}
  {{else if targetHasRoleName $usertag "Hunter"}}
    {{takeRoleName $usertag "Hunter"}}
    {{sleep 30}}
    {{giveRolename $usertag "Hunter"}}
  {{else if targetHasRoleName $usertag "Druid"}}
    {{takeRoleName $usertag "Druid"}}
    {{sleep 30}}
    {{giveRolename $usertag "Druid"}}
  {{end}}
  
{{end}}

