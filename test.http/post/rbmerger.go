package post

func (rbMerger *RBMerger) postMergedRBs(descName string) error {
	var rbList []appv1alpha1.ResourceBinding
	desc, err := rbMerger.localGaiaClient.AppsV1alpha1().Descriptions(common.GaiaReservedNamespace).Get(context.TODO(), descName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	rbs, err := rbMerger.rbLister.ResourceBindings(common.GaiaRBMergedReservedNamespace).List(labels.SelectorFromSet(labels.Set{
		common.GaiaDescriptionLabel: descName,
		common.StatusScheduler:      string(appv1alpha1.ResourceBindingmerged),
	}))
	if err != nil {
		return err
	}
	for _, rb := range rbs {
		rbList = append(rbList, *rb)
	}
	resultSchemaSet := &resourcebindingmerger.SchemaSet{
		AppID:  descName,
		RBList: rbList,
		Desc:   *desc,
	}
	postBody, err := json.Marshal(resultSchemaSet)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", rbMerger.postURL, bytes.NewReader(postBody))
	if err != nil {
		klog.Infof("PostHyperOM: ERROR: failed create post request %q.", err)
		return err
	}
	req.Close = true
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}
	klog.Infof("PostHyperOM: Begin post description %q to HyperOM.")

	resp, err := httpClient.Do(req)
	defer func() { err = resp.Body.Close() }()
	if err != nil || resp.StatusCode != http.StatusOK {
		klog.Infof("PostHyperOM: ERROR: post to HyperOM error %q.", err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		klog.Infof("PostHyperOM: ERROR: response from HyperOM error %q.", err)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		klog.Infof("PostHyperOM: ERROR: post to HyperOM error %q.", err)
		return err
	}

	return nil
}
