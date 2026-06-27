export const DataView = {
  // Giao diện template dạng chuỗi chứa vùng xem dữ liệu và form chỉnh sửa
  template: `
<div :class="theme.dataview">
  <div :class="theme.dataview_header"> {{ title }} </div>
</div>
  `,
  props: {
    title: {
      type: String,
      required: true
    },
    labels: {
      type: Object,
      required: true
    },
    datas: {
      type: Object,
      required: true
    }
  },
  emits: ['update:datas'],
  setup(props, { emit }) {
    const { ref, reactive, watch } = Vue;

    // Trạng thái chuyển đổi giữa Xem và Chỉnh sửa
    const isEditing = ref(false);
    
    // Bản ghi dữ liệu đang hiển thị hiện tại
    const currentData = ref({ ...props.datas });
    
    // Vùng nhớ đệm lưu trữ dữ liệu trên Form nhằm tránh thay đổi trực tiếp dữ liệu gốc khi chưa nhấn Lưu
    const formData = reactive({});

    // Theo dõi thay đổi từ prop 'datas' bên ngoài để đồng bộ khi cần thiết
    watch(() => props.datas, (newData) => {
      currentData.value = { ...newData };
    }, { deep: true });

    // Khởi động chế độ chỉnh sửa: Đổ dữ liệu hiện tại vào Form
    const startEdit = () => {
      Object.assign(formData, currentData.value);
      isEditing.value = true;
    };

    // Lưu lại thông tin từ Form sang dữ liệu hiển thị và phát tín hiệu báo cho lớp cha
    const saveEdit = () => {
      currentData.value = { ...formData };
      isEditing.value = false;
      // Gửi dữ liệu mới ra ứng dụng cha theo chuẩn dữ liệu hai chiều
      emit('update:datas', { ...currentData.value });
    };

    // Hủy bỏ trạng thái chỉnh sửa
    const cancelEdit = () => {
      isEditing.value = false;
    };

    return {
      isEditing,
      currentData,
      formData,
      startEdit,
      saveEdit,
      cancelEdit
    };
  }
};
