{{exec "sb" "stealth"}}
{{$msg:= joinStr "" .User.Username " suddenly disappears!"}} 
{{sendMessage nil $msg}}
{{giveRoleName .User "Stealth"}}
