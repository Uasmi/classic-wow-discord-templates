{{$usertag:= userArg (index .CmdArgs 0)}}
{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
Oops, you're on cooldown. 
{{else}}
{{$msg:= joinStr "" "Holy spirit guards you, " $usertag.Username "!"}} 
  {{sendMessage nil $msg}}
  {{sendMessage nil ":sparkles:"}}
  {{giveRoleName $usertag "Bubble"}}
  {{takeRoleName $usertag "Bubble" (100)}}
  {{dbSetExpire .User.ID ("usedCC") ("True") (100)}}
{{end}}
