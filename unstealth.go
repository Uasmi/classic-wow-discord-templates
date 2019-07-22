{{$msg:= joinStr "" "Wild " .User.Username " appears"}} 
{{sendMessage nil $msg}}
{{takeRoleName .User "Stealth"}}
