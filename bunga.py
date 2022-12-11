import turtle


t = turtle.Turtle()


jumlah_petal = 10


t.color("purple")
t.width(3)


for i in range(jumlah_petal):
    t.forward(100)
    t.right(45)
    t.forward(100)
    t.right(135)

turtle.done()