package services

// func (uc *CsiUseCase) push(d []byte) {
// pack := uc.decoder.DecodeCsiPackage(d)

// if pack.Info.CsiLength == 0 {
// 	return
// }

// if uc.isFilterActive && !uc.filter.Check(pack.Info) {
// 	return
// }

// pack.Uuid = uuid.New().String()
// pack.Timestamp = time.Now().UnixMilli()
// pack.Number = uc.csiPackageNumber
// uc.csiPackageNumber += 1

// log.Print("PUSH ", uc.csiPackageNumber)

// uc.repo.Push(pack)

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

// apiPack := entity.ApiPackageAbsPhase{
// 	Timestamp: pack.Timestamp,
// 	Id:        pack.Uuid,
// 	Info:      pack.Info,
// 	Number:    pack.Number,
// 	Abs:       abs,
// 	Phase:     uc.proc.CsiMap(pack.Data, processor.PhaseHandler),
// }

// uc.cbPushPacket(apiPack)
// }

// func (uc *CsiUseCase) OnPushPacket(cb func(entity.ApiPackageAbsPhase)) {
// 	uc.cbPushPacket = cb
// }
