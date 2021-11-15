document.getElementById("hover").addEventListener("mouseenter", function () {
  document.getElementById("draw").classList.add("animate");
  
  let audio = document.getElementById("music");
  if (audio.paused){
     audio.play();
  }
});

document.getElementById("hover").addEventListener("mouseleave", function () {
  document.getElementById("draw").classList.remove("animate");
});
