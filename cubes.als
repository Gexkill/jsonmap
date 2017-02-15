abstract sig Cube {
onTop : lone Cube
}

sig Large, Medium, Small extends Cube {}

sig TopCube in Large + Medium + Small {}

fact {
// reflexivity is not valid
--all c:Cube | c.onTop != c

// simmetry is not valid
--all disj c,c':Cube | c.onTop = c' implies c'.onTop != c

// simmetry and reflexivity are not valid
all c,c':Cube | c.onTop = c' implies c'.onTop != c

// transitivity is not valid
all disj c,c',c'':Cube | c.onTop = c' and  c''.onTop = c implies  c'.onTop != c''

// 2 cube cannot be onTop another cube
all disj c,c',c'':Cube | not (c.onTop = c' and c''.onTop = c' )

/*
Large can be onTop only another Large Cube
Medium can be onTop another Medium or Large Cube
Small can be onTop another Small, Medium or Large Cube
*/
all c,c':Cube | c.onTop = c' implies ((c in Small and c' in (Small + Medium + Large)) 
or (c in Medium and c' in (Medium + Large)) 
or (c in Large and c' in (Large)))

all disj c,c':TopCube | not (c.onTop = c' or c'.onTop = c)
all disj c,c':Cube | c in TopCube implies c.onTop = c' and all c'':Cube | not (c''.onTop = c)
}

pred show [] {}

pred putOnTop (A,B:Cube,A',B':Cube){

// precondition
A.onTop = none

//postcondition
B'.onTop = B.onTop
A'.onTop = B'

}

--run putOnTop

run show
