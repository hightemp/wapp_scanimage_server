<div class="main-wrapper">
    <div class="top-panel">Сканирование документов</div>
    <div class="middle-wrapper">
        <div class="left-side">
            <div style="overflow-y:scroll;">
                <?php echo $sScannersList ?>
            </div>
            <form class="left-side-form" action="?m=Main&a=fnActionHTML" target="actions-iframe" method="POST">
                <div>
                    <button type="sumbit" class="btn btn-success" name="action" value="scan">Скан.</button>
                    <button type="sumbit" class="btn btn-danger" name="action" value="delete">Удал.</button>
                    <button type="sumbit" class="btn btn-danger" name="action" value="delete_all">Очис.</button>
                    <button type="sumbit" class="btn btn-info" name="action" value="download">Скач.</button>
                    <button type="sumbit" class="btn btn-info" name="action" value="convert_pdf">PDF</button>
                    <button type="button" class="btn btn-primary" name="action" value="refresh" onlcik="window.location.reload()">Обн.</button>
                </div>
                <div style="overflow-y:scroll;">
                    <div class="list-group">
                        <?php foreach ($aScannedFiles as $aScannedFile): ?>
                        <div class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <input type="checkbox" value="<?php echo $aScannedFile[0] ?>" name="images[]" /> 
                                <a href="?m=Main&a=fnImageHTML&file=<?php echo $aScannedFile[0] ?>" target="image-iframe"><h5 class="mb-1"><?php echo $aScannedFile[0] ?></h5></a>
                            </div>
                            <p class="mb-1"><?php echo $aScannedFile[2] ?></p>
                        </div>
                        <?php endforeach ?>
                    </div>
                </div>

                <nav>
                    <div class="nav nav-tabs" id="nav-tab" role="tablist">
                        <a class="nav-item nav-link active" id="nav-home-tab" data-toggle="tab" href="#nav-home" role="tab" aria-controls="nav-home" aria-selected="true">Архивы</a>
                        <a class="nav-item nav-link" id="nav-profile-tab" data-toggle="tab" href="#nav-profile" role="tab" aria-controls="nav-profile" aria-selected="false">PDF</a>
                    </div>
                </nav>
                <div class="tab-content" id="nav-tabContent">
                    <div class="tab-pane fade show active" id="nav-home" role="tabpanel" aria-labelledby="nav-home-tab">
                        <div>
                            <button type="sumbit" class="btn btn-danger" name="action" value="delete_archives">Удал.</button>
                            <button type="sumbit" class="btn btn-danger" name="action" value="delete_all_archives">Очис.</button>
                        </div>
                        <div style="overflow-y:scroll;">
                            <div class="list-group">
                                <?php foreach ($aPackedFiles as $aPackedFile): ?>
                                <div class="list-group-item list-group-item-action flex-column align-items-start">
                                    <div class="d-flex w-100 justify-content-between">
                                        <input type="checkbox" value="<?php echo $aPackedFile[0] ?>" name="archives[]" /> 
                                        <a href="<?php echo $aPackedFile[1] ?>" target="_blank"><h5 class="mb-1"><?php echo $aPackedFile[0] ?></h5></a>
                                    </div>
                                    <p class="mb-1"><?php echo $aPackedFile[2] ?></p>
                                </div>
                                <?php endforeach ?>
                            </div>
                        </div>
                    </div>
                    <div class="tab-pane fade" id="nav-profile" role="tabpanel" aria-labelledby="nav-profile-tab">
                        <div>
                            <button type="sumbit" class="btn btn-danger" name="action" value="delete_pdf">Удал.</button>
                            <button type="sumbit" class="btn btn-danger" name="action" value="delete_all_pdf">Очис.</button>
                        </div>
                        <div style="overflow-y:scroll;">
                            <div class="list-group">
                                <?php foreach ($aPDFFiles as $aPDFFile): ?>
                                <div class="list-group-item list-group-item-action flex-column align-items-start">
                                    <div class="d-flex w-100 justify-content-between">
                                        <input type="checkbox" value="<?php echo $aPDFFile[0] ?>" name="pdf[]" /> 
                                        <a href="<?php echo $aPDFFile[1] ?>" target="_blank"><h5 class="mb-1"><?php echo $aPDFFile[0] ?></h5></a>
                                    </div>
                                    <p class="mb-1"><?php echo $aPDFFile[2] ?></p>
                                </div>
                                <?php endforeach ?>
                            </div>
                        </div>
                    </div>
                </div>


            </form>
        </div>
        <div class="right-side">
            <iframe src="<?php echo $sURL ?>" frameborder="0" id="actions-iframe" name="actions-iframe" style="display:none"></iframe>

            <iframe src="?m=Main&a=fnImageHTML" frameborder="0" id="image-iframe" name="image-iframe"></iframe>
        </div>
    </div>
</div>
