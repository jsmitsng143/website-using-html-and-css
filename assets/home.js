var image1=new Image()
image1.src="./images/slides/1.png"
var image2=new Image()
image2.src="./images/slides/2.png"
var image3=new Image()
image3.src="./images/slides/3.png"

var step=1
function slideit(){
document.images.slide.src=eval("image"+step+".src")
if(step<3)
step++
else
step=1
setTimeout("slideit()",2500)
}
slideit()