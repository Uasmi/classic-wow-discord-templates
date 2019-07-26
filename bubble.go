{{$usertag:= userArg (index .CmdArgs 0)}}
{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
Oops, you're on cooldown. 
 {{else if targetHasRoleName $usertag "Banished"}}
    {{sendMessage nil "He's at the other realm... :smiling_imp:"}}
{{else}}
 {{exec "sb" "bubble2"}}
{{$msg:= joinStr "" "Holy spirit guards you, " $usertag.Username "!"}} 
  {{sendMessage nil $msg}}
  {{sendMessage nil ":sparkles:"}}
  {{giveRoleName $usertag "Bubble"}}
  {{takeRoleName $usertag "Bubble" (150)}}
  {{dbSetExpire .User.ID ("usedCC") ("True") (10)}}
{{end}}
