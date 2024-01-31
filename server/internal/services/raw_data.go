package services

// Сглаживание
// Нужно эту обработку вынести в WebSocketController
// там будет вытягиваться отдельный сервис обработки
// abs := uc.proc.CsiMap(pack.Data, processor.AbsHandler)

// if uc.csiPackageNumber > uint64(uc.smoothOrder) {
// 	prevs := uc.repo.GetLastN(uc.smoothOrder)

// 	for i := 0; i < uc.smoothOrder; i++ {
// 		prev_abs := uc.proc.CsiMap(prevs[i].Data, processor.AbsHandler)
// 		for j := 0; j < 4; j++ {
// 			for k := 0; k < 56; k++ {
// 				abs[j][k] += prev_abs[j][k]
// 			}
// 		}
// 	}

// 	for j := 0; j < 4; j++ {
// 		for k := 0; k < 56; k++ {
// 			abs[j][k] /= float64(uc.smoothOrder)
// 		}
// 	}
// }
// Конец сглаживания
