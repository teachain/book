<template>
    <Card>
        <Form :model="formItem">
            <FormItem label="挖掘机名称" :label-width="80" label-position="top">
                <Input v-model="formItem.name" placeholder="请输入挖掘机名称..." />
            </FormItem>
            <FormItem label="挖掘机型号" :label-width="80" label-position="top">
                <Select v-model="formItem.model_id">
                    <Option v-for="item in models" :value="item.id" :key="item.id">{{ item.name }}</Option>
                </Select>
            </FormItem>
            <FormItem :label-width="0">
                <Button type="primary" v-on:click="onSave">保存</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    data() {
        return {
            formItem: {
                name: "",
                model_id: 0,
            },
            models: [],
        }
    },
    created() {
        this.$http.get("/model/list").then((data) => {
            this.models = data.data.data;
        });
    },
    methods: {
        onSave() {
            if (this.formItem.name == "") {
                this.$Message.error('挖掘机名称不能为空');
                return;
            }
            if (this.formItem.model_id == 0) {
                this.$Message.error('请选择挖掘机型号');
                return;
            }
            this.$http.post('/device/add', { name: this.formItem.name, model_id: this.formItem.model_id }).then((data) => {
                if (data.data.code == 0) {
                    this.$Message.success("挖掘机信息保存成功");
                    this.$router.push({ path: "/" });
                } else {
                    this.$Message.error("挖掘机信息保存失败");
                }
            });
        }
    }
}
</script>