# SongJSON Specification
SongJSON (probably going to change the name) is a JSON subset that is suitable for song lyrics, chords, notes, and metadata. It is inspired by the OpenSong XML format. Each JSON string represents a single song.

## Object keys

### `title`
*Required*
The song title as a string.

### `body`
*Required*
The `body` value is string representation of the body of the song. Each line in the string begins with a character that denotes what type of information the line contains.

```javascript
" ", or "" // A line beginning with a space or any other character contains lyrics
";" // A line beginning with a semi-colon contains notes
"!" // A line beginning with a exclamation point denotes a section change
"." // A line beginning with a period contains chords
```

If a lyric line appears before a section line, this line is interpreted as being the beginning of a new anonymous section. If there is a non-blank `section` value in the surrounding object, anonymous sections are ignored.

For lines with chords to be valid, they must be followed by a line with lyrics that begins with a space. Otherwise they are ignored. A chord and lyric line couplet "line up," that is, the chord is at the same position on its line as the lyrics it goes with on the line below it.

If there is a need to put more chords in a space on a section of lyrics than could fit typically, underscores are used as a filler on the lyric line.

If there is a line that contains nothing or only white-space and is surrounded by lyric lines, that is interpreted as the beginning of a new page in the section. If the lyrics contain no sections, anonymous sections are derived from these lines.

**Example 1**
```json
{
    "body": ";Key of B
!v1
.           G              D                 C        D        G
 Come, Thou Fount of every blessing, Tune my heart to sing Thy grace;
.           Em           D                 C        D       G
 Streams of mercy, never ceasing, Call for songs of loudest praise.
.         Em     C       G               Em      C        D
 Teach me some melodious sonnet, Sung by flaming tongues above;
.           G                 D                C     D       G
 Praise the mount, I'm fixed upon it, Mount of Thy redeeming love.",
    "etc": "..."
}
```
This example starts with a note denoting the key, A section, `v1`, and a number of chord lines followed by lyric lines.

**Example 2**
```json
{
    "body": "Twinkle, twinkle, little star,
How I wonder what you are.
Up above the world so high,
Like a diamond in the sky.
Twinkle, twinkle, little star,
How I wonder what you are!

When the blazing sun is gone,
When there's nothing he shines upon,
Then you show your little light,
Twinkle, twinkle, through the night.
Twinkle, twinkle, little star,
How I wonder what you are!",
    "etc": "..."
}
```
This example contains only lyrics and doesn't have a section line. But the lack of any beginning section and the blank line makes this song have two anonymous sections.

### `presentation`
*Optional*
The `presentation` value represents the order of the song by space-delimited sections. The sections are referenced in the `body` value. Section values are case-insensitive.

There are particular section types that also represent a common song sections. These are represented by the implementation as their long name. They are as follows:

```javascript
"c" // Representation: Chorus
"b" // Representation: Bridge
"v#" // Representation: Verse #
// Verses are "v" followed by any number, which is the verse number
"pc" // Representation: Pre-chorus
"i" // Representation: Intro
"e" // Representation: Ending
```

Section types that don't fit these patterns are represented as just their original string.

**Example**
```json
{
    "presentation": "v1 c v2 c b Interlude c",
    "etc": "..."
}
```
The `body` value must then contain each the sections `v1`, `c`, `v2`, `b`, and `Interlude` (can also be included as `interlude` as it is case-insensitive). These are interpreted as `Verse 1`, `Chorus`, `Verse 2`, `Bridge`, and `Interlude` respectively.

### `author`
*Optional*
The song author

### `ccli`
*Optional*
The Christian Copyright Licensing International number