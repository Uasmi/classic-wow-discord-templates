{{$usertag:= .ExecData.ID}}
{{$time:= .ExecData.Time}}
{{if targetHasRoleName $usertag "Mage"}}
      {{takeRoleName $usertag "Mage"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Mage"}}
    {{else if targetHasRoleName $usertag "Warrior"}}
      {{takeRoleName $usertag "Warrior"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Warrior"}}
    {{else if targetHasRoleName $usertag "Warlock"}}
      {{takeRoleName $usertag "Warlock"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Warlock"}}
    {{else if targetHasRoleName $usertag "Shaman"}}
      {{takeRoleName $usertag "Shaman"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Shaman"}}
    {{else if targetHasRoleName $usertag "Rogue"}}
      {{takeRoleName $usertag "Rogue"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Rogue"}}
    {{else if targetHasRoleName $usertag "Priest"}}
      {{takeRoleName $usertag "Priest"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Priest"}}
    {{else if targetHasRoleName $usertag "Paladin"}}
      {{takeRoleName $usertag "Paladin"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Paladin"}}
    {{else if targetHasRoleName $usertag "Hunter"}}
      {{takeRoleName $usertag "Hunter"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Hunter"}}
    {{else if targetHasRoleName $usertag "Druid"}}
      {{takeRoleName $usertag "Druid"}}
      {{sleep $time}}
      {{giveRoleName $usertag "Druid"}}
    {{end}}
