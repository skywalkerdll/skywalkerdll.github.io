
function playmusic(){
  let audio = document.getElementById("music");
  if (audio.paused){
     audio.play();
  }
}

window.onload = function () {
  document.getElementById("hover").addEventListener("mouseenter", function () {
    document.getElementById("draw").classList.add("animate");
    setInterval(playmusic(),1000); 
  
  });
  
  document.getElementById("hover").addEventListener("mouseleave", function () {
    document.getElementById("draw").classList.remove("animate");
  });

}
