package repository

import "testing"

func BenchmarkAdminRepository_FindByName(b *testing.B) {

	adminRepo := &AdminRepository{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = adminRepo.FindByName("admin")
	}
	b.StopTimer()
}
