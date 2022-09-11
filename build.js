const Fs = require('fs')
const { exec } = require('child_process')
const JSZip = require('jszip');

const images = Fs.readdirSync("./source")
const cmd = 
        `magick montage ${images.map(x => `"./source/${x}"`).join(" ")} -shadow 5 -geometry 32x32+2+2 ./build/face.png`
exec(cmd, (err, stdout, stderr) => {
        if (err) {
                console.error(err)
        }
})

const zip = new JSZip();
for (let i of images) {
        zip.file(i, Fs.readFileSync(`./source/${i}`))
}
zip.file("README.md", Fs.readFileSync(`README.md`))

zip.generateNodeStream({ type: 'nodebuffer', streamFiles: true })
        .pipe(Fs.createWriteStream('./build/RPGCharactersPack.zip'))
        .on('finish', function () {
            console.log("Done");
        });
