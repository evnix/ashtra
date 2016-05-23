package server
import ("testing")


func TestCreate(t *testing.T){


	CreateMetaFile("testq",10)

	qfp:= QFileOp{};

	qfp.OpenMetaFile("testq")

}