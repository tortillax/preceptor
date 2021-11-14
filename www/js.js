var playlistContentData = "";
var playlistData = "";
var channelsData = "";

async function selectPlaylist(id) {
    fetch("/action/setPlaylist/" + id);

    await new Promise(r => setTimeout(r, 550));
    buttonReload();
}

function selectServer(server, channel) {
    fetch("/action/connect/" + server + "/" + channel);
}

function buttonDisconnect() {
    fetch("/action/disconnect");
}

function buttonPlay() {
    fetch("/action/play");
}

function buttonStop() {
    fetch("/action/stop");
}

function buttonNext() {
    fetch("/action/next");
}

function buttonReload() {
    //status
    var selStatus = document.getElementById("status");
    fetch("/info/status")
    .then(response => response.json())
    .then(function(data) {
        if (data.status === "idle") {
            selStatus.innerHTML = "<b>status</b>: idle";
        } else {
            selStatus.innerHTML = "<b>status</b>: playing "+data.status;
        }
    })
    .catch(function(error) {        
        selStatus.innerHTML = "<b>status</b>: error comunicating with API";
    });

    //playlist
    var selPlaylistContent = document.getElementById("playlistContent");
    fetch("/info/playlist")
    .then(response => response.json())
    .then(function(data) {   
        playlistContentData = "";  
        for (i = 0; i < data.length; i++) {
            playlistContentData += "<tr><td>" + data[i] + "</td></tr>";
        }
        
        selPlaylistContent.innerHTML = playlistContentData;
    })
    .catch(function(error) {
        playlistContentData += "<tr><td>error comunicating with API</td></tr>";        
        selPlaylistContent.innerHTML = playlistContentData;
    });

    //playlists
    var selPlaylists = document.getElementById("playlists");
    fetch("/info/playlists")
    .then(response => response.json())
    .then(function(data) {
        playlistData = "";
        for (i = 0; i < data.length; i++) {
            playlistData += "<button class='u-full-width' onClick='selectPlaylist(" + i + ")'>" + data[i] + "</button>";
        }
        selPlaylists.innerHTML = playlistData;
    })
    .catch(function(error) {
        playlistData = "error comunicating with API";
        selPlaylists.innerHTML = playlistData;
    });    

    //channels
    var selChannels = document.getElementById("channels");
    fetch("/info/allChannels")
    .then(response => response.json())
    .then(function(data) {
        channelsData = "";
        for (var i = 0; i < data.length; i++) {
            channelsData += "<button class='u-full-width' onClick='selectServer(\"" + data[i].ServerID + "\", \"" + data[i].ChannelID + "\")'>" + data[i].ChannelName + " @" + data[i].ServerName + "</button>";
        }
        selChannels.innerHTML = channelsData;
    })
    .catch(function(error) {
        channelsData = "error comunicating with API";
        selChannels.innerHTML = channelsData;
    });
}

/*
    fetch("/info/status")
    .then(response => response.json())
    .then(function(data) {

    })
    .catch(function(error) {
        
    })
    */