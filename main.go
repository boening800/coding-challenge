package main

import "github.com/gofiber/fiber/v3"
import "gonum.org/v1/gonum/mat"

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type QRResponse struct {
	Q[][]float64 `json:"Q"`
	R[][]float64 `json:"R"`
}

func handleRefactorQR(c fiber.Ctx) error {
	var req MatrixRequest

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid json",
			})
		}

		rows := len(req.Matrix)
		cols := len(req.Matrix[0])

		data := []float64{}

		for i := 0; i < rows; i++ {
			data = append(data, req.Matrix[i]...)
		}

		A := mat.NewDense(rows, cols, data)

		var qr mat.QR
		qr.Factorize(A)

		var Q mat.Dense
		var R mat.Dense

		qr.QTo(&Q)
		qr.RTo(&R)

		qRows, qCols := Q.Dims()
		rRows, rCols := R.Dims()

		qMatrix := make([][]float64, qRows)
		rMatrix := make([][]float64, rRows)

		for i := 0; i < qRows; i++ {
			qMatrix[i] = make([]float64, qCols)
			for j := 0; j < qCols; j++ {
				qMatrix[i][j] = Q.At(i, j)
			}
		}

		for i := 0; i < rRows; i++ {
			rMatrix[i] = make([]float64, rCols)
			for j := 0; j < rCols; j++ {
				rMatrix[i][j] = R.At(i, j)
			}
		}

		return c.JSON(QRResponse{
			Q: qMatrix,
			R: rMatrix,
		})
}

func main() {
	
	app :=fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/refactor-qr", handleRefactorQR)

	app.Listen(":3000")

}