document.addEventListener('DOMContentLoaded', () => {
    const getTableName=() => {
        const path = window.location.pathname;
        const parts = path.split('/');
        return parts[parts.length - 1]; // Lấy phần cuối cùng của đường dẫn
    }
    // Chỉ hiển thị nút Thêm 
    document.querySelectorAll('.add-btn').forEach(button => {
        button.addEventListener('click', () => {
            window.location.href = `/${getTableName()}/add`; // Chuyển hướng tới trang chỉnh sửa
        });
    });
    // Xử lý nút Sửa
    document.querySelectorAll('.edit-btn').forEach(button => {
        button.addEventListener('click', () => {
            const id = button.getAttribute('data-id');
            window.location.href = `/${getTableName()}/edit?id=${id}`; // Chuyển hướng tới trang chỉnh sửa
        });
    });

    // Xử lý nút Xóa
    document.querySelectorAll('.delete-btn').forEach(button => {
        button.addEventListener('click', async () => {
            const id = button.getAttribute('data-id');
            if (!confirm('Bạn có chắc muốn xóa mục này?')) {
                return;
            }

            try {
                const response = await fetch(`/${getTableName()}/delete?id=${id}`, {
                    method: 'DELETE',
                });

                if (response.ok) {
                    alert('Xóa thành công!');
                    window.location.reload(); // Tải lại trang
                } else {
                    const error = await response.text();
                    alert(`Lỗi: ${error}`);
                }
            } catch (err) {
                alert('lỗi kết nối server');
            }
        });
    });
});