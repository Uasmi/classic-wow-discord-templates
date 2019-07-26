{{$cmd:= (exec "sb" "stealth")}}
{{/* $cmd */}}
{{$msg:= joinStr "" .User.Username " suddenly disappears!"}} 
{{sendMessage nil $msg}}
{{giveRoleName .User "Stealth"}}
  
