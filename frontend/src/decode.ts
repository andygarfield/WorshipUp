interface SongJSON {
    title: string
    body: string
    presentation?: string
    author?: string
    ccli?: string
}

export function decodeSong(sj: SongJSON) {
    // Return a blank string if there is no song data yet
    if (!sj) {
        return ""
    }
    // Otherwise, render the data

    // Variable holding the HTML data to be returned
    let outHtml = ""

    // Helper function for quickly adding elements to the rendered output
    let appendElement = (tagName: string, className: string | null, innerText: string) => {
        outHtml += `
            <${tagName}${className ? ` class="` + className + `"`: ``}>${innerText}</${tagName}>
        `;
    }

    // Add author
    if (sj.author) {
        appendElement('h3', 'author', sj.author)
    }

    
    let chordLine = "";
    let pendingLineBreak = false;

    if (!sj.body) {
        appendElement('div', 'comment', "This song has no content")
    }
    // Use the first character of every line to identify its line-type
    for (let line of sj.body.split("\n")) {
        let firstChar = line.slice(0, 1)

        // If there is a seperation between the lyrics which isn't before a new
        // section, insert a <br> tag
        if (pendingLineBreak && firstChar != "!") {
            outHtml += '<br>';
            pendingLineBreak = false;
        } else if (pendingLineBreak && firstChar == "!") {
            pendingLineBreak = false;
        }

        if (line) {
            switch (line[0]) {
                //Comment line
                case ';':
                    appendElement('div', 'comment', line.slice(1).trim())
                    break;
                // Chord line
                case '.':
                    // Save the line for use with a lyric line
                    chordLine = line.slice(1);
                    break;
                // Lyric line
                case ' ':
                    if (chordLine) {
                        appendElement('div', 'couplet-line', parseChords(chordLine, line.slice(1)))
                        chordLine = "";
                    } else {
                        appendElement('div', 'lyric', line.slice(1));
                    }
                    break;
                // Section line
                case '!':
                    let secondChar = line.slice(1, 2).toLowerCase();
                    let expanded = '';
    
                    // See if the section fits a known type
                    // The bug here is that any section that begins
                    // with these letters will be assumed to be these
                    // sections. There also needs to be an implimentation
                    // of numbered choruses and bridges and the like
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
                    appendElement('section', 'song-section', expanded)
                    break;
            }
        } else {
            pendingLineBreak = true;
        }

    }
    return outHtml;
}

function parseChords(chordLine: string, bodyLine: string) {
    let chordLineArray = chordLine.split("")

    let chordIndices = [];
    let currentChord = "";

    // Loop through the chord characters
    for (let i = 0; i < chordLineArray.length; i++) {
        // If there is no current chord
        if (currentChord == "") {
            // Check if current is the beginning of a chord
            if (isChordLetter(chordLineArray[i])) {
                currentChord += chordLineArray[i]
            }
        // If there is a current chord
        } else {
            // If a space is the next character, the chord has ended
            if (chordLineArray[i] == " ") {
                chordIndices.push([i - currentChord.length, i])
                currentChord = ""
            // If not, the chord is added to
            } else {
                currentChord += chordLineArray[i]
            }
        }
    }
    // Catch the chords at the end of the line
    if (currentChord != "") {
        chordIndices.push([chordLineArray.length - currentChord.length, chordLineArray.length])
    }

    // If the lyric line is shorter than the chord line,
    // normalize it to be as long as chord line
    if (chordLine.length > bodyLine.length) {
        let difference = chordLine.length - bodyLine.length
        for (let i = 0; i < difference; i++) {
            bodyLine += " "
        }
    }

    // An array for the pairs of chords and lyrics
    let clPairs = [];
    for (let i = 0; i < chordIndices.length; i++) {
        // If it's the first time through the loop...
        if (i == 0) {
            // ... and if there's a beginning (where the 
            // first chord character is not index 0)...
            if (chordIndices[0][0] != 0) {
                // ... write a beginning
                clPairs.push([
                    '&nbsp;',
                    bodyLine.slice(0, chordIndices[0][0])
                        .replace(/^\s+|\s+$/g, '&nbsp;'),
                ])
            } // Otherwise, continue
        } else {
            clPairs.push([
                chordLine.slice(chordIndices[i - 1][0], chordIndices[i - 1][1]),
                bodyLine.slice(chordIndices[i - 1][0], chordIndices[i][0])
                    .replace(/^\s+|\s+$/g, '&nbsp;'),
            ])
        }
    }
    // Add final value
    clPairs.push([
        chordLine.slice(
            chordIndices[chordIndices.length - 1][0], chordIndices[chordIndices.length - 1][1]
        ),
        bodyLine.slice(chordIndices[chordIndices.length - 1][0])
            .replace(/^\s+|\s+$/g, '&nbsp;'),
    ])

    // Loop through pairs and make divs
    let newDiv = ""
    for (let clPair of clPairs) {
        newDiv += `
            <div class="c-l-couplet">
                <div class="chord">${clPair[0]}</div>
                <div class="lyric">${clPair[1]}</div>
            </div>
        `
    }
    return newDiv
}

function isChordLetter(char: string) {
    let chordLetters = ['a', 'b', 'c', 'd', 'e', 'f', 'g']
    if (chordLetters.indexOf(char.toLowerCase()) != -1) {
        return true
    }
    return false
}