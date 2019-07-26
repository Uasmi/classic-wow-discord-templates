{{$usertag:= userArg (index .CmdArgs 0)}}
{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
Oops, you're on cooldown. 
{{else}}
  {{if targetHasRoleName $usertag "Bubble"}}
    {{$msg:= joinStr "" $usertag.Username " is guarded by the Holy Spirit :sparkles:!"}}
    {{sendMessage nil $msg}}
  {{else}}
    {{exec "sb" "sheep"}}
    {{$msg:= joinStr "" "You were polymorphed " $usertag.Username "!"}} 
    {{sendMessage nil $msg}}
    {{sendMessage nil ":sheep:Baaah:sheep: "}}
    {{giveRoleName $usertag "Sheep"}}
    {{takeRoleName $usertag "Sheep" (30)}}
    {{dbSetExpire .User.ID ("usedCC") ("True") (10)}}
  {{end}}
{{end}}
