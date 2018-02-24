let randomString = Math.random().toString(36).substring(7);

function setup() {
  canvas = createCanvas(100, 100);
  noStroke();
  // fill(255, 0, 0);
  // rect(0, 0, 33, height);
  // fill(0, 255, 0);
  // rect(33, 0, 66, height);
  // fill(0, 0, 255);
  // rect(66, 0, width, height);
}

function draw() {
  background(255);
  textSize(15)
  text(frameCount, width/2, height/2);
  sendCanvas(canvas);
  if (frameCount % 25 == 0) {
    console.log(`fc: ${frameCount} fr: ${frameRate()}`)
  }
  if (frameCount > 300) {
    noLoop();
  }
}

function sendCanvas(imageData) {
  let frameNum = padDigits(frameCount, 4);
  let filename = `abc_${randomString}_${frameNum}.png`;
  // let payload = {
  //   filename, 
  //   "canvas": imageData.canvas.toDataURL(),
  // };
  let echoUrl = "http://localhost:3246";
  let url = "http://localhost:8080/image";
  let formData = new FormData();
  formData.append('filename', filename);
  formData.append('canvas', imageData.canvas.toDataURL())
  fetch(url, {
    // headers: {
    //   'Accept': 'application/json, application/xml, text/plain, text/html, *.*',
    //   'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
    // },
    body: formData,
    method: 'POST',
  }).then(response => {
    // debugger;
    console.log(response.text());
  });
}


function padDigits(number, digits) {
  return Array(Math.max(digits - String(number).length + 1, 0)).join(0) + number;
}