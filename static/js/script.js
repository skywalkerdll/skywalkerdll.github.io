<<<<<<< HEAD
=======
document.getElementById("hover").addEventListener("mouseenter", function () {
  document.getElementById("draw").classList.add("animate");
  setInterval(playmusic(),1000); 
});

document.getElementById("hover").addEventListener("mouseleave", function () {
  document.getElementById("draw").classList.remove("animate");
});
 
>>>>>>> 83e55dca729a4cc651f47887a913b356db63deba

function playmusic(){
  let audio = document.getElementById("music");
  if (audio.paused){
     audio.play();
  }
<<<<<<< HEAD
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
=======
}
>>>>>>> 83e55dca729a4cc651f47887a913b356db63deba
