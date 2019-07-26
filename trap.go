{{$record:= dbGet (.User.ID) ("usedCC")}}

{{if in $record.Value "True"}}
Oops, you're on cooldown. 
{{else}}
  {{dbSet 0 ("trap") (.User.Username)}}
  {{sendDM "Tss, your trap is activated.."}}
  {{dbSetExpire .User.ID ("usedCC") ("True") (10)}}
  {{deleteMessage nil .Message.ID (1)}}
{{end}}
