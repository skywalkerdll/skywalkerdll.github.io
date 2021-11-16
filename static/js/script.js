document.getElementById("hover").addEventListener("mouseenter", function () {
  document.getElementById("draw").classList.add("animate");
  setInterval(playmusic(),1000); 

});

document.getElementById("hover").addEventListener("mouseleave", function () {
  document.getElementById("draw").classList.remove("animate");
});

// var svgdocment=document.getElementById('svgobject').contentDocument;
// var svgdocment = document.getElementById('svgembed').getSVGDocument();

// svgdocment.getElementById("hover").addEventListener("mouseenter", function () {
//   svgdocment.getElementById("draw").classList.add("animate");
//   setInterval(playmusic(),1000); 

// });

// svgdocment.getElementById("hover").addEventListener("mouseleave", function () {
//   svgdocment.getElementById("draw").classList.remove("animate");
// });

function playmusic(){
  let audio = document.getElementById("music");
  if (audio.paused){
     audio.play();
  }
}