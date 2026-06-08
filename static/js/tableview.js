const TableView ={
  template: `
  <div :class="theme.tv_main">
      <div :class="theme.tv_header">
        <div :class="theme.tv_title"> {{ title }} </div>
        <div :class="theme.tv_subtitle"> {{ subtitle }} </div>
      </div>
      <div :class="theme.tv_body">
        <table :class="theme.tv_table">
          <thead :class="theme.tv_table_header">
            <tr>
              <th v-for="label in labels" :key="label.key" > {{ label.label }} </th>
              <th v-if="showAction">Thao tác</th>
            </tr>
          </thead>
          <tbody :class="theme.tv_table_body">
            <tr v-for="(item,index) in datas" :key="index">
              <td v-for="label in labels" :key="label.key" > {{ item[label.key] }} </td>
              <td v-if="showAction" :class="theme.tv_table_cell_edit_button" >
                <button :class="[theme.tv_button, theme.tv_button_edit]" @click="openEditModal(index)">
                  <span :class="theme.tv_button_text">Sửa</span>
                </button>
                <button :class="[theme.tv_button, theme.tv_button_delete]" @click="deleteItem(index)">
                  <span :class="theme.tv_button_text">Xóa</span>
                </button>
              </td>
            </tr>
          </tbody>
          <tfoot :class="theme.tv_table_footer">
            <tr>
              <td colspan="100%" style="padding: 8px;">
                <div :class="theme.tv_pagination">
                  <span :class="theme.tv_pagination_text">Tổng số: {{ datas.length }} </span>
                  <button :class="theme.tv_pagination_button" @click="openAddModal">Thêm mới</button>
                </div>
              </td>
            </tr>
          </tfoot>
        </table>
      </div>
      <div :class="theme.tv_footer">
        <div :class="theme.tv_footer_text"> {{ footer }} </div>
      </div>
  </div>
  `,
  // ... (Giữ nguyên phần props, emits và setup của bạn)
  props: {
    title: { type: String, default: 'Danh sách' },
    subtitle: { type: String, default: '' },
    footer: { type: String, default: '' },
    labels: { type: Array, default: () => [] },
    datas: { type: Array, default: () => [] },
    showAction: { type: Boolean, default: false },
  },
  emits: ['addItem', 'updateItem', 'deleteItem'],
  setup(props, context){
    const openEditModal = (index) => {
      context.emit('updateItem', props.datas[index]);
    };
    const openAddModal = () => {
      context.emit('addItem');
    };
    const deleteItem = (index) => {
      if(confirm('Bạn có chắc muốn xóa?'))
        context.emit('deleteItem', index);
    };
    return {
      theme: AppTheme,
      openEditModal,
      openAddModal,
      deleteItem,
    };
  },
};
