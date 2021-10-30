# Terrible Audio Downloader

This is just a small go project I did for myself, to manage my audio library to get away from spotify.
All it does is download music specified in a json.
The songs are seperated by gerne and each genre makes a folder.
It doesn't support playlist, but a workaround would be to make the playlist a genre.
To listen to the music, use a audioplayer like mocp or vlc. 

## Dependencies
- youtube-dl

## Usage
There are 2 arguments which must be set.
-p sets the path, where it'll download the music too.
Keep in mind, that it'll create a subfolder for each genre.
-f sets the file, where you configured your music.
The structure is explained in the next section.

## Config
The configuration is just the file you provide, where you defined the music you want to download.
The structure is as follows:

```json
{
	"genres": {
		"<Your Genre": {
			"<SongName>": {
				"artists": [
					"<Artist1>",
					"<Artist2>"
				],
				"links": [
					"https://www.youtube.com/watch?v=dQw4w9WgXcQ"
				]
			},
			"<SongName2": {
				...
			}
		},
		"<Second Genre>": {
			...
		}
	}
}
```

## TODO
- Download Music concurrent
- Remove Songs not specified
