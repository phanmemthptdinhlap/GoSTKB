// DynamicTable.js
import { ref } from 'vue';

export default {
  name: 'DynamicTable',
  props: {
    headers: {
      type: Array,
      required: true
    },
    modelValue: {
      type: Array,
      required: true
    },
    title: {
      type: String,
      default: 'Danh sách dữ liệu'
    }
  },
  emits: ['update:modelValue'],

  setup(props, { emit }) {
    // Khai báo các trạng thái phản xạ (Reactive State) đúng chuẩn Vue 3
    const editingItem = ref(null);
    const editingIndex = ref(-1);

    // Các hàm xử lý (thay thế cho methods)
    const startEdit = (item, index) => {
      editingIndex.value = index;
      editingItem.value = { ...item }; // Sao chép nông để tránh thay đổi trực tiếp
    };

    const saveChanges = () => {
      if (editingIndex.value !== -1) {
        // Tạo mảng mới từ props và cập nhật phần tử sửa đổi
        const updatedList = [...props.modelValue];
        updatedList[editingIndex.value] = editingItem.value;

        // Phát sự kiện cập nhật v-model về component cha
        emit('update:modelValue', updatedList);
        cancelEdit();
      }
    };

    const cancelEdit = () => {
      editingItem.value = null;
      editingIndex.value = -1;
    };

    // Trả về các biến/hàm để sử dụng ngoài template
    return {
      editingItem,
      editingIndex,
      startEdit,
      saveChanges,
      cancelEdit
    };
  },

  template: `
    <div class="dynamic-table-box" style="margin-bottom: 40px; padding: 20px; border: 1px solid #ddd; border-radius: 8px;">
      <h2>{{ title }}</h2>
      
      <table border="1" cellpadding="8" cellspacing="0" style="width: 100%; border-collapse: collapse; margin-bottom: 20px;">
        <thead>
          <tr style="background-color: #f2f2f2;">
            <th v-for="header in headers" :key="header.key">{{ header.label }}</th>
            <th style="width: 100px;">Hành động</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in modelValue" :key="index">
            <td v-for="header in headers" :key="header.key">
              {{ item[header.key] }}
            </td>
            <td align="center">
              <button @click="startEdit(item, index)">Sửa</button>
            </td>
          </tr>
          <tr v-if="modelValue.length === 0">
            <td :colspan="headers.length + 1" align="center">Không có dữ liệu</td>
          </tr>
        </tbody>
      </table>

      <div v-if="editingItem" class="edit-form-container" style="background: #f9f9f9; padding: 15px; border-radius: 4px;">
        <h3>Chỉnh sửa dòng số {{ editingIndex + 1 }}</h3>
        <div v-for="header in headers" :key="'form-' + header.key" style="margin-bottom: 10px;">
          <label style="display: inline-block; width: 120px; font-weight: bold;">{{ header.label }}:</label>
          <input 
            v-model="editingItem[header.key]" 
            :disabled="header.key === 'id'"
            style="padding: 5px; width: 250px;"
          />
        </div>
        <div style="margin-top: 15px; margin-left: 120px;">
          <button @click="saveChanges" style="background: #4CAF50; color: white; border: none; padding: 6px 12px; cursor: pointer; margin-right: 10px;">Lưu</button>
          <button @click="cancelEdit" style="background: #f44336; color: white; border: none; padding: 6px 12px; cursor: pointer;">Hủy</button>
        </div>
      </div>
    </div>
  `
};
