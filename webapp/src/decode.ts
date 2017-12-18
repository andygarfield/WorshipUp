interface SongJSON {
    title: string
    lyrics: string
    presentation?: string
    author?: string
    ccli?: string
}

function parseChords(chordLine: string, lyricsLine: string) {
    let chordLineArray = chordLine.split("")
    let chordIndices = [];

    let currentChord = "";
    for (let i = 0; i < chordLineArray.length; i++) {
        if (currentChord == "") {
            if (isChordLetter(chordLineArray[i])) {
                currentChord += chordLineArray[i]
            } else {
                new Error("Not a chord")
            }
        } else {
            if (chordLineArray[i] == " ") {
                chordIndices.push([i - currentChord.length, i])
                currentChord = ""
            } else {
                currentChord += chordLineArray[i]
            }
        }
    }

    let newDiv = ""
    for (let chordLineIndex of chordIndices) {
        if (newDiv == "") {
            let beginning = lyricsLine.slice(0, chordLineIndex[0])
            if (beginning.length != 0) {
                newDiv += `
                    <div class="cl-couplet">
                        <div class="chords"></div>
                        <div class="lyrics">${beginning}</div>
                    </div>
                    `
            }
            newDiv += `
            <div class="cl-couplet">
                <div class="chords">${chordLine.slice(chordLineIndex[0], chordLineIndex[1])}</div>
                <div class="lyrics">${lyricsLine.slice(chordLineIndex[0], chordLineIndex[1])}</div>
            </div>
            `
        } else {
            // newDiv += newMiddleCouplet(chordLineIndex)
        }
    }

    return newDiv

    // function newMiddleCouplet(indices: number[]) {
    //     return `
    //     <div class="cl-couplet">
    //         <div class="chords">${chordLine.slice(indices[0], indices[1])}</div>
    //         <div class="lyrics">${lyricsLine.slice(indices[0], indices[1])}</div>
    //     </div>
    //     `
    // }

    function isChordLetter (char: string) {
        let chordLetters = ['a', 'b', 'c', 'd', 'e', 'f', 'g']
        if (chordLetters.indexOf(char.toLowerCase()) != -1) {
            return true
        }
        return false
    }
}

export default function decode(sj: SongJSON) {
    let outHtml = ""
    var previousLineTag = "";

    let appendElement = (tagName: string, className: string | null, innerText: string) => {
        if (previousLineTag == tagName) {
            outHtml += "<br>" + innerText
        } else {
            outHtml += `</${previousLineTag}><${tagName}${className ? ` class="` + className + `"`: ""}>${innerText}`;
        }
        previousLineTag = tagName;
    }

    // Add title and author
    appendElement('h1', 'title' , sj.title)
    if (sj.author) {
        appendElement('h3', 'author', sj.author)
    }

    let chordLine = "";
    for (let line of sj.lyrics.split("\n")) {
        let firstChar = line.slice(0, 1)
        switch (line[0]) {
            case ';':
                appendElement('div', 'comment', line.slice(1).trim())
                break;
            case '.':
                chordLine = line.slice(1);
                break;
            case ' ':
                if (chordLine) {
                    appendElement('div', 'couplet-line', parseChords(chordLine, line.slice(1)))
                    chordLine = "";
                } else {
                    appendElement('div', 'lyrics', line.slice(1));
                }
                break;
            case '!':
                let secondChar = line.slice(1, 2).toLowerCase();
                let expanded = '';

                switch (secondChar) {
                    case "v":
                        expanded = "Verse " + line.slice(2);
                        break;
                    case "c":
                        expanded = "Chorus";
                        break;
                    case "b":
                        expanded = "Bridge";
                        break;
                    case "p":
                        expanded = "Pre-Chorus";
                        break;
                    case "i":
                        expanded = "Intro";
                        break;
                    case "e":
                        expanded = "Ending";
                        break;
                    default:
                        expanded = line.slice(1);
                }
                appendElement('section', 'songSection', expanded)
                break;
        }
    }
    return outHtml;
}