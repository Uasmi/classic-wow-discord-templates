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
    {{sendMessage nil "You're bringing him back. Pathetic"}}
    {{takeRoleName $usertag "Banished"}}
  {{else}}
    {{$cmd:=(exec "sb" "banish")}}
    {{/* $cmd */}}
    {{$msg:= joinStr "" $usertag.Username " was banished! :smiling_imp: !"}} 
    {{sendMessage nil $msg}}
    {{giveRoleName $usertag "Banished"}}
    {{takeRoleName $usertag "Banished" (30)}}
    {{dbSetExpire .User.ID ("usedCC") ("True") (2)}}
    {{execCC 8 nil 0 (sdict "ID" $usertag.ID "Time" 30)}}
  {{end}}
{{end}}
