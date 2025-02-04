ssh cynxhost@cynx.buzz "rm /home/cynxhost/cynxhost-server/cynxhost"
scp cynxhost cynxhost@cynx.buzz:/home/cynxhost/cynxhost-server/cynxhost
ssh cynxhost@cynx.buzz "service cynxhost-server restart"