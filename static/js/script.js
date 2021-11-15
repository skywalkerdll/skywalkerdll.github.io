document.getElementById("hover").addEventListener("mouseenter", function () {
  document.getElementById("draw").classList.add("animate");
});

document.getElementById("hover").addEventListener("mouseleave", function () {
  document.getElementById("draw").classList.remove("animate");
});

let audio = document.getElementById("music");
audio.play();
