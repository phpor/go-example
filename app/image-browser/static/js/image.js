let img = document.getElementById("image")
let images = [];
let i = 0

window.onload = function () {
    for (let ele of document.getElementById("images").children) {
        images.push(ele.textContent)
    }
    if (images.length > 0) {
        show_image(0)
    }
}

function show_image(index) {
    img.src = images[index]
    img.style.display = "block"
}

// 在按下向左、向右箭头时触发一个函数
window.addEventListener(
    "keydown",
    (event) => {
        if (event.key == "ArrowLeft") {
            i--;
            if (i < 0 ) i = 0
        }
        if (event.key == "ArrowRight") {
            i++
            if (i >= images.length ) i = images.length - 1
        }

        show_image(i)
    },
    true,
);