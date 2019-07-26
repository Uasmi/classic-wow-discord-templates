{{$args:= parseArgs 1 "You have to mention someone!"
  (carg "user" "usr")}}
{{$usertag:=$args.Get 0}}
{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
Oops, you're on cooldown. 
{{else}}
  {{exec "sb" "dispel"}}
  {{$msg:= joinStr "" "You're purified, " $usertag.Username "!"}} 
  {{sendMessage nil $msg}}
  {{sendMessage nil ":star2:"}}
  {{if targetHasRoleName $usertag "Sheep"}}
      {{takeRoleName $usertag "Sheep"}}
  {{else if targetHasRoleName $usertag "Banished"}}
      {{sendMessage nil "He's at the other realm... :smiling_imp:"}}
   {{else if targetHasRoleName $usertag "Stealth"}}
      Hmmm, did you really see someone?
   {{end}}
   {{dbSetExpire .User.ID ("usedCC") ("True") (10)}}
{{end}}
