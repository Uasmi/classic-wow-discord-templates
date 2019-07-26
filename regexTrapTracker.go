{{$trap:= dbGet 0 ("trap")}}
{{if in $trap.Value "False"}}
{{else}}
  {{if in ($trap.Value) .User.Username}}
  {{else}}
  You activated a trap!
  {{giveRoleName .User.ID "Frozen"}}
  {{takeRoleName .User.ID "Frozen" (15)}}
  {{dbSet 0 ("trap") ("False")}}
  {{execCC 8 nil 0 (sdict "ID" .User.ID "Time" 15)}}
  {{if targetHasRoleName .User.ID "Stealth"}}
  {{takeRoleName .User.ID "Stealth"}}
  {{end}}
  {{end}}
{{end}}
