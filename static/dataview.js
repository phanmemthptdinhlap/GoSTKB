const DataView = {
  template: `
  <div :class="theme.dataview">
    <div :class="theme.dataview_header">
      <div :class="theme.dataview_title">{{ title }}</div>
    </div>
    <div :class="theme.dataview_body">
      <div :class="theme.dataview_item" v-for="(item,index) in datas" :key="index">
        <div v-for ="(label,index) in labels" :key="label.key">
          <div :class="theme.dataview_item_label">{{ label.label }}</div>
          <div :class="theme.dataview_item_value">{{ item[label.key] }}</div>
        </div>
        <div :class="theme.dataview_action">
          <button :class="[theme.dataview_button_edit]" @click="openEditForm(item)">
            Chỉnh sửa
          </button>
          <button :class="[theme.dataview_button_delete]" @click="deleteData(item)">
            Xóa
          </button>
        </div>
      </div>
      <div :class="theme.dataview_footer">
        <button :class="[theme.dataview_button theme.dataview_button_add]" @click="openAddForm">
          Thêm mới 
        </button>
        <button :class="[theme.dataview_button theme.dataview_button_refresh]" @click="syncData"> 
          Đồng bộ dữ liệu
        </button>
      </div>
    </div>
  </div>
  <div :class="theme.formview" v-if="showForm" >
    <!-- FORM -->
    <div :class="theme.formview_header">
      <div :class="theme.formview_title">{{ formTitle }}</div>
    </div>

    <div :class="theme.formview_body">
      <div :class="theme.formview_form">
        <div v-for="(label,index) in labels" :key="label.key" :class="theme.formview_item" >
          <label :class="theme.formview_item_label" > {{ label.label }}: </label>
          <input :class="theme.formview_item_input" type="text" v-model="data[label.key]">
        </div>
      </div>
      <div :class="theme.formview_button">
        <button :class="[theme.formview_button_save]" @click="saveData" :disabled="isLoading">
          {{ isLoading ? 'Đang lưu ... ' : 'Lưu' }}
        </button>
        <button :class="[theme.formview_button_cancel]" @click="closeForm" :disabled="isLoading">
          Hủy
        </button>
      </div>
    </div>
    
  </div>
`,

  props : {
    title: { type: String, default: 'Danh sách' },
    data: { type: Array, default: () => [] },
    labels: { type: Array, default: () => [] },
  },
  emits : ['syncData, deleteData'],

  setup(props, context){
    const {ref} = Vue;
    const {emit, expose} = context;
    // ... (Giữ nguyên khai báo biến và hàm của bạn) ...
    const showForm = ref(false);
    const isLoading = ref(false);
    const formTitle = ref('Thêm mới');
    const data = ref({});
      
    const closeForm = () => { showForm.value = false; isLoading.value = false; data.value = {}; };
    const openEditForm = (dataEdit) => { showForm.value = true; isLoading.value = false; formTitle.value = 'Chỉnh sửa'; data.value = {...dataEdit}; };
    const openAddForm = () => { showForm.value = true; isLoading.value = false; formTitle.value = 'Thêm mới'; data.value = {}; };
    const saveData = () => {
      isLoading.value = true;
      emit('saveData', data.value,(res) => {
        isLoading.value = false;
        if(res.success){ closeForm(); } else { message.value = res.message; }
        });
      };

    expose({ openEditForm, openAddForm, closeForm });
    
    return {
      theme: AppTheme, // QUAN TRỌNG: Phải return theme ở đây để template dùng được
      showForm,
      isLoading,
      formTitle,
      data,
      closeForm,
      saveData,
    };
  }
}; 

      
