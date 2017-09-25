#declare ShortNeck = cylinder {
    <0,-1,0>, <0,0.6,0>, 0.2
}

#declare LongNeck = cylinder {
    <0,-1,0>, <0,0.8,0>, 0.2
}

#declare Base = merge {
    cylinder {
        <0,-1.1,0>, <0,-1,0>, 0.6
    }
    cylinder {
        <0,-1,0>, <0,-0.9,0>, 0.5
    }
}

#declare ShortBody = merge {
    object { Base }
    object { ShortNeck }
}

#declare LongBody = merge {
    object { Base }
    object { LongNeck }
}
