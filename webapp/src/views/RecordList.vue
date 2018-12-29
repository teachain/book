<template>
    <div>
        <Card>
            <Form>
                <FormItem label="名称" :label-width="80">
                </FormItem>
                <FormItem label="型号" :label-width="80">
                </FormItem>
                <FormItem label="" :label-width="80">
                </FormItem>
                <FormItem label="" :label-width="80">
                </FormItem>
                <FormItem label="" :label-width="80">
                </FormItem>
                <FormItem label="" :label-width="80">
                </FormItem>
                <FormItem label="" :label-width="80">
                </FormItem>
                <FormItem :label-width="0">
                </FormItem>
            </Form>
        </Card>
        <Button type="primary" v-on:click="goBack()">返回</Button>
    </div>
</template>
<script>
export default {
    data() {
        return {
            records: [],
        }
    },
    created() {
        this.$http.get("/record/list").then((data) => {
            this.records = data.data.data;
        });
    },
    methods: {
        goBack() {
            this.$router.push({ path: "/" });
        }
    }
}
</script>