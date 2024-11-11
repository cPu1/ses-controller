func asdf() {
    res := &svcsdk.GetTemplateInput{}
    {{ GoCodeSetCreateInput .CRD "r.ko.Spec" "res" 1 }}
}