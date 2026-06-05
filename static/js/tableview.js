const TableView ={
  template: `
  <div class="tv-main">
      <div class="tv-header">
        <div class="tv-title"> {{ title }} </div>
        <div class="tv-subtitle"> {{ subtitle }} </div>
      </div>
      <div class="tv-body">
        <table class="tv-table">
          <thead class="tv-table-header">
            <tr>
              <th v-for="label in labels" :key="label.key" > {{ label.label }} </th>
              <th v-if="showAction">Thao tác</th>
            </tr>
          </thead>
          <tbody class="tv-table-body">
            <tr v-for="(item,index) in datas" :key="index">
              <td v-for="label in labels" :key="label.key" > {{ item[label.key] }} </td>
              <td v-if="showAction" class="tv-table-cell-edit-button" >
                <button class="tv-button tv-button-edit" @click="openEditModal(index)">
                  <span class="tv-button-text">Sửa</span>
                </button>
                <button class="tv-button tv-button-delete" @click="deleteItem(index)">
                  <span class="tv-button-text">Xóa</span>
                </button>
              </td>
            </tr>
          </tbody>
          <tfoot class="tv-table-footer">
            <tr>
              <td colspan="100%" style="padding: 8px;">
                <div class="tv-pagination">
                  <span class="tv-pagination-text">Tổng số: {{ datas.length }} </span>
                  <button class="tv-pagination-button" @click="openAddModal">Thêm mới</button>
                </div>
              </td>
            </tr>
          </tfoot>
        </table>
      </div>
      <div class="tv-footer">
        <div class="tv-footer-text"> {{ footer }} </div>
      </div>
  </div>
  `,
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
      openEditModal,
      openAddModal,
      deleteItem,
    };
  },
};

