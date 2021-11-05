# preceptor - discord music bot
 - play local mp3 files on any discord channel

## API endpoints

| endpoint                               | description                           |
| -------------------------------------- | ------------------------------------- |
| /info/playlists                        | list all available playlists          |
| /info/playlist                         | list songs in current playlist        |
| /info/status                           | get status info (song)                |
| /info/servers                          | list servers                          |
| /info/channels/*serverID*              | list voice channels in server         |
| /action/connect/*serverID*/*channelID* | conenct to channel on specific server |
| /action/disconnect                     | disconnect from server                |
| /action/play                           | start (or resume) playing             |
| /action/next                           | play next song                        |
| /action/stop                           | stop playing                          |
| /action/setPlaylist/*id*               | set and randomize playslist           |