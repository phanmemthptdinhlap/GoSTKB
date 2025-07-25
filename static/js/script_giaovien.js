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
                    <button class="button" onclick="suaGiaoVien('${giaovien.ma_giao_vien}','${giaovien.ho_ten}','${giaovien.ten_tkb}')">Sửa</button>
                    <button class="button" onclick="xoaGiaoVien('${giaovien.ma_giao_vien}')">Xóa</button>
                </td>
                </tr>`;
                bang.innerHTML+=row;
            });
        document.getElementById("title").textContent="Danh sách giáo viên";
        document.getElementById("divform").style.display="none";
        });
}
function xoaGiaoVien(ma_giao_vien){
    fetch(`/api/giaovien/${ma_giao_vien}`,{method:'DELETE'})
    .then(()=>LayDanhSachGiaoVien());
}
//Sửa giáo viên
function suaGiaoVien(ma_giao_vien,ho_ten,ten_tkb){
    document.getElementById("title").textContent="Sửa giáo viên"
    document.getElementById("divform").style.display="block";
    document.getElementById("form_edit").value="edit";
    document.getElementById("ma_giao_vien").value=ma_giao_vien;
    document.getElementById("ho_ten").value=ho_ten;
    document.getElementById("ten_tkb").value=ten_tkb;
}
//Thêm giáo viên vào danh sách
document.getElementById("btThemMoi").addEventListener("click",function(e){
    e.preventDefault();
    document.getElementById("title").textContent="Thêm mới giáo viên"
    document.getElementById("divform").style.display="block";
    document.getElementById("form_edit").value="";
});
document.getElementById("FormGiaoVien").addEventListener("submit",function(e){
    e.preventDefault();
    const formedit=document.getElementById("form_edit").value;
    const giaovien={
        ma_giao_vien:document.getElementById("ma_giao_vien").value,
        ho_ten:document.getElementById("ho_ten").value,
        ten_tkb:document.getElementById("ten_tkb").value
    };
    fetch(formedit?`/api/giaovien/${giaovien.ma_giao_vien}`:`/api/giaovien`,{
        method:formedit?"PUT":"POST",
        headers:{'Content-Type':'application/json'},
        body:JSON.stringify(giaovien)
    }).then(()=>{
        LayDanhSachGiaoVien();
        document.getElementById('FormGiaoVien').reset();
    });
});

//Tải danh sách giáo viên
document.addEventListener('DOMContentLoaded',LayDanhSachGiaoVien);