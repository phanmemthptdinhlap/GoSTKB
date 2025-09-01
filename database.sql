CREATE TABLE IF NOT EXISTS giaovien (
		ma_giao_vien INTEGER PRIMARY KEY AUTOINCREMENT,
		ho_ten TEXT NOT NULL,
		ten_tkb TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS monhoc (
		ma_mon INTEGER PRIMARY KEY AUTOINCREMENT,
		ten_mon TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS lophoc (
		ma_lop INTEGER PRIMARY KEY AUTOINCREMENT,
		ten_lop TEXT NOT NULL,
		khoi_lop TEXT NOT NULL,
		ma_chu_nhiem INTEGER REFERENCES giaovien(ma_giao_vien)
);
CREATE TABLE IF NOT EXISTS phancong (
		ma_phan_cong INTEGER PRIMARY KEY AUTOINCREMENT,
		ma_giao_vien INTEGER REFERENCES giaovien(ma_giao_vien),
		ma_mon INTEGER REFERENCES monhoc(ma_mon_hoc),
		ma_lop INTEGER REFERENCES lophoc(ma_lop_hoc)
);
CREATE TABLE IF NOT EXISTS tietday(
		ma_tiet_day INTEGER PRIMARY KEY	AUTOINCREMENT,
		ma_phan_cong INTEGER REFERENCES phancong(ma_phan_cong),
		tuan INTEGER NOT NULL,
		tong_tiet_duoc_phan_cong INTEGER NOT NULL,
		tiet_sang INTEGER,
		tiet_chieu INTEGER,
		tiet_no INTEGER
);