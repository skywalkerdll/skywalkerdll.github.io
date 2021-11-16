// document.getElementById("hover").addEventListener("mouseenter", function () {
//   document.getElementById("draw").classList.add("animate");
//   setInterval(playmusic(),1000); 

// });

// document.getElementById("hover").addEventListener("mouseleave", function () {
//   document.getElementById("draw").classList.remove("animate");
// });

window.onload = function () {

  var a=document.getElementById('svgObject');
  var svgdocment=a.contentDocument;
  //var svgdocment = document.getElementById('svgembed').getSVGDocument();

  svgdocment.getElementById("hover").addEventListener("mouseenter", function () {
    document.getElementById("draw").classList.add("animate");
    setInterval(playmusic(),1000); 

  });

  svgdocment.getElementById("hover").addEventListener("mouseleave", function () {
    document.getElementById("draw").classList.remove("animate");
  });

}

function playmusic(){
  let audio = document.getElementById("music");
  if (audio.paused){
     audio.play();
  }
}