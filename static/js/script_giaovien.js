function LayDanhSachGiaoVien(){
    fetch('/api/giaovien')
        .then(response=>response.json())
        .then(data=>{
            const bang=document.getElementById("BangGiaoVien");
            bang.innerHTML='';
            data.forEach(giaovien=>{
                const row=`<tr>
                <td>${giaovien.ma_giao_vien}</td>
                <td>${giaovien.ho_ten}</td>
                <td>${giaovien.ten_tkb}</td>
                <td>
                    <button class="edit" onclick="suaGiaoVien(${giaovien.ma_giao_vien})">Sửa</button>
                    <button class="delete" onclick="xoaGiaoVien(${giaovien.ma_giao_vien})">Xóa</button>
                </td>
                </tr>`;
                bang.innerHTML+=row;
            });
        document.getElementById("divform").className="hiden";
        document.getElementById("divview").className="show";
        });
}
function xoaGiaoVien(id){
    fetch(`/api/giaovien/${id}`,{method:'DELETE'})
    .then(()=>LayDanhSachGiaoVien());
}
//Thêm giáo viên vào danh sách
document.getElementById("btThemMoi").addEventListener("click",function(e){
    e.preventDefault();
    document.getElementById("title").textContent="Thêm mới giáo viên"
    document.getElementById("divform").className="show";
    document.getElementById("divview").className="hiden";
});
document.getElementById("FormGiaoVien").addEventListener("submit",function(e){
    e.preventDefault();
    const giaovien={
        ma_giao_vien:document.getElementById("ma_giao_vien").value,
        ho_ten:document.getElementById("ho_ten").value,
        ten_tkb:document.getElementById("ten_tkb").value
    };
    fetch('/api/giaovien',{
        method:'POST',
        headers:{'Content-Type':'application/json'},
        body:JSON.stringify(giaovien)
    }).then(()=>{
        LayDanhSachGiaoVien();
        document.getElementById('FormGiaoVien').reset();
    });
});

//Tải danh sách giáo viên
document.addEventListener('DOMContentLoaded',LayDanhSachGiaoVien);