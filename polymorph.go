{{$args:= parseArgs 1 "You have to mention someone!"
  (carg "user" "usr")}}
{{$usertag:=$args.Get 0}}
{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
Oops, you're on cooldown. 
{{else}}
  {{if targetHasRoleName $usertag "Bubble"}}
    {{$msg:= joinStr "" $usertag.Username " is guarded by the Holy Spirit :sparkles:!"}}
    {{sendMessage nil $msg}}
  {{else if targetHasRoleName $usertag "Stealth"}}
    {{sendMessage nil "Hmmm, did you really see someone?"}}
  {{else if targetHasRoleName $usertag "Banished"}}
    {{sendMessage nil "He's at the other realm... :smiling_imp:"}}
  {{else}}
    {{$cmd:=(exec "sb" "sheep")}}
    {{/* $cmd */}}
    {{$msg:= joinStr "" "You were polymorphed " $usertag.Username "!"}} 
    {{sendMessage nil $msg}}
    {{sendMessage nil ":sheep:Baaah:sheep: "}}
    {{giveRoleName $usertag "Sheep"}}
    {{takeRoleName $usertag "Sheep" (15)}}
    {{dbSetExpire .User.ID ("usedCC") ("True") (10)}}
    {{execCC 8 nil 0 (sdict "ID" $usertag.ID "Time" 15)}}
  {{end}}
{{end}}
