mbmatch
=======

An MBTiles server for PBF.

Pass it an openmaptiles mbtiles file datasource, it can serve the map without an internet connection.

Pass environment:

`HOSTNAME` `TILESPATH` `PORT`

or command flags:
```
Usage of ./mbmatch:
  -debug
        enable debug
  -hostname string
        the hostname to come back at tiles (default "localhost:7000")
  -port int
        port to listen for HTTP (default 7000)
  -tilesPath string
        mbtiles file path
```

Open your browser at http://localhost:7000

It's mainly use for an [offline car map project](https://blog.nobugware.com/post/2018/my_own_car_system_raspberry_pi_offline_mapping/) but can also serve tiles on the web.

## Installation

You need the Go compiler installed and set a GOPATH env variable.

```
go get github.com/akhenakh/mbmatch/cmd/mbmatch
${GOPATH}/bin/mbmatch -path mymap.mbtiles -hostname myserver:8000
```

or use [a docker image](https://cloud.docker.com/repository/docker/akhenakh/mbmatch).

```
docker run -it --rm -p 8000:8000 -e HOSTNAME=mymap.server.com -e TILESPATH=/root/hawaii.mbtiles akhenakh/mbmatch:latest  
```

## Autostart
Edit `/etc/systemd/system/mbmatch.service`

```
[Unit]
Description=mbmatch

[Service]
WorkingDirectory=/home/youruser
ExecStart=/home/youruser/mbmatch -tilesPath /home/youruser/north-america.mbtiles 
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```
`systemctl enable mbmatch`

## License 

Code is licensed under MIT.

It contains several assets:

The Mapbox GL Style JSON file is originally derived from [OSM Bright from Mapbox Open Styles](https://github.com/mapbox/mapbox-gl-styles/blob/master/LICENSE.md). The modified Mapbox GL Style JSON retains the same BSD license.

> Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

> * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
* Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
* Neither the name of Mapbox nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

### Design License

The look and feel of the OSM liberty map design is also derived (although heavily altered) from [OSM Bright from Mapbox Open Styles](https://github.com/mapbox/mapbox-gl-styles/blob/master/LICENSE.md) which is licensed under the Creative Commons Attribution 3.0 license.

### Icons

OSM Liberty is using the [Maki POI icon set](https://github.com/mapbox/maki/blob/master/LICENSE.txt) which is licensed under CC0 1.0 Universal.

### Fonts

OSM Liberty is using the Roboto font family (Copyright 2011 Google).
