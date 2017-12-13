interface SongJSON {
    title: string
    lyrics: string
    presentation?: string
    author?: string
    ccli?: string
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

    
    for (let line of sj.lyrics.split("\n")) {
        let firstChar = line.slice(0, 1)
        switch (line[0]) {
            case ';':
                appendElement('div', 'comment', line.slice(1).trim())
                break;
            case '.':
                appendElement('div', 'chords', line.slice(1))
                break;
            case ' ':
                appendElement('div', 'lyrics', line.slice(1).trim())
                break;
            case '!':
                let secondChar = line.slice(1, 2).toLowerCase()
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