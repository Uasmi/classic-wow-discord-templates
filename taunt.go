{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
Oops, you're on cooldown. 
{{else}}
 {{$msg:= (exec "sb" "taunt")}}
  {{/* $msg */}}
  {{dbSetExpire .User.ID ("usedCC") ("True") (10)}}
{{end}}
